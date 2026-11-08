import assert from 'node:assert/strict';
import { mkdtemp, readdir, readFile, rm, stat, writeFile } from 'node:fs/promises';
import { tmpdir } from 'node:os';
import { resolve, join } from 'node:path';
import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';

const cacheDir = await mkdtemp(join(tmpdir(), 'floyd-supercache-test-'));
const environment = Object.fromEntries(
  Object.entries(process.env).filter(([, value]) => typeof value === 'string'),
);
environment.FLOYD_SUPERCACHE_DIR = cacheDir;

const transport = new StdioClientTransport({
  command: process.execPath,
  args: [resolve('dist/index.js')],
  env: environment,
  stderr: 'pipe',
});
const client = new Client({ name: 'supercache-integration-test', version: '1.0.0' });

function decode(result) {
  const text = result.content?.find((item) => item.type === 'text')?.text;
  assert.equal(typeof text, 'string');
  return JSON.parse(text);
}

async function call(name, args) {
  return decode(await client.callTool({ name, arguments: args }));
}

try {
  await client.connect(transport);

  const tools = await client.listTools();
  assert.equal(tools.tools.length, 12);

  const resources = await client.listResources();
  assert.deepEqual(
    resources.resources.map(({ name }) => name).sort(),
    ['health', 'tool-registry'],
  );
  const health = await client.readResource({ uri: 'mcp://floyd-supercache-server/health.json' });
  assert.equal(JSON.parse(health.contents[0].text).toolCount, 12);

  assert.equal((await call('cache_store', { key: 'alpha/beta', value: { id: 1 } })).success, true);
  assert.equal((await call('cache_store', { key: 'alpha:beta', value: { id: 2 } })).success, true);
  assert.deepEqual((await call('cache_retrieve', { key: 'alpha/beta' })).value, { id: 1 });
  assert.deepEqual((await call('cache_retrieve', { key: 'alpha:beta' })).value, { id: 2 });

  const projectFiles = await readdir(join(cacheDir, 'project'));
  assert.equal(projectFiles.length, 2, 'distinct keys must not collide on disk');
  for (const file of projectFiles) {
    assert.equal((await stat(join(cacheDir, 'project', file))).mode & 0o777, 0o600);
  }
  assert.equal((await stat(cacheDir)).mode & 0o777, 0o700);
  assert.equal((await stat(join(cacheDir, 'index.json'))).mode & 0o777, 0o600);

  const wildcard = await call('cache_list', { filter: 'alpha*' });
  assert.equal(wildcard.total, 2);
  const literalPattern = await call('cache_list', { filter: 'alpha[beta' });
  assert.equal(literalPattern.total, 0, 'filter punctuation must be treated literally');
  assert.equal((await call('cache_search', { query: 'alpha' })).found, 2);

  await Promise.all(Array.from({ length: 20 }, (_, index) => call('cache_store', {
    key: `concurrent-${index}`,
    value: index,
  })));
  assert.equal((await call('cache_list', { filter: 'concurrent-*' })).total, 20);

  assert.equal((await call('cache_store_reasoning', {
    context: 'integration-test',
    reasoning: 'validate persistence',
    conclusion: 'works',
  })).tier, 'reasoning');
  assert.equal((await call('cache_load_reasoning', { context: 'integration-test' })).conclusion, 'works');
  assert.equal((await call('cache_archive_reasoning', { context: 'integration-test' })).success, true);

  assert.equal((await call('cache_clear', { tier: 'all', confirm: false })).error, 'Confirmation required');
  assert.equal((await call('cache_clear', { tier: 'all', confirm: true })).success, true);
  assert.equal((await call('cache_stats', { tier: 'all' })).totalEntries, 0);

  const index = JSON.parse(await readFile(join(cacheDir, 'index.json'), 'utf8'));
  assert.deepEqual(index.entries, {});

  await writeFile(join(cacheDir, 'index.json'), '{not-json', 'utf8');
  const corruptResult = await client.callTool({ name: 'cache_stats', arguments: { tier: 'all' } });
  assert.equal(corruptResult.isError, true, 'corrupt indexes must fail closed');
  assert.match(decode(corruptResult).error, /JSON/);

  console.log('Supercache integration tests: startup, storage, concurrency, permissions, resources, archival, and fail-closed corruption passed');
} finally {
  await client.close().catch(() => undefined);
  await rm(cacheDir, { recursive: true, force: true });
}
