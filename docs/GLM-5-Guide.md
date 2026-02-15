# GLM-5: Technical Reference for Coding Agents

> **Source:** Official Z.ai Documentation (https://docs.z.ai, https://docs.bigmodel.cn)

---

## Model Overview

**GLM-5** is Z.ai's new-generation flagship foundation model, designed for **Agentic Engineering**. It provides reliable productivity in complex system engineering and long-range Agent tasks.

**Key Characteristics:**
- Achieves state-of-the-art (SOTA) performance in open source for Coding and Agent capabilities
- Real programming scenario usability approaching Claude Opus 4.5
- Ideal foundation model for general Agent assistants

---

## API Configuration

### Endpoints

| Platform | Base URL |
|----------|----------|
| Z.ai International | `https://api.z.ai/api/paas/v4/` |
| BigModel.cn (China) | `https://open.bigmodel.cn/api/paas/v4/` |

### Model Name

```
glm-5
```

### Authentication

```bash
Authorization: Bearer YOUR_API_KEY
```

---

## API Parameters

### Official Documentation Examples

The following parameters are demonstrated in official Z.ai documentation:

| Parameter | Type | Description |
|-----------|------|-------------|
| `model` | string | `"glm-5"` |
| `messages` | array | Chat message history |
| `thinking` | object | Thinking mode configuration |
| `stream` | boolean | Enable streaming output |
| `max_tokens` | integer | Maximum output tokens |
| `temperature` | float | Output randomness control |

### Thinking Mode

The `thinking` parameter controls extended reasoning:

```json
{
  "thinking": {
    "type": "enabled"
  }
}
```

- **Values:** `"enabled"` | `"disabled"`
- **Default:** `"enabled"` (from official docs)

### Streaming Response Structure

When streaming with thinking enabled, responses include:

```json
{
  "choices": [{
    "delta": {
      "reasoning_content": "<reasoning process>",
      "content": "<final response>"
    }
  }]
}
```

---

## Code Examples

### cURL - Basic Call

```bash
curl -X POST "https://api.z.ai/api/paas/v4/chat/completions" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your-api-key" \
-d '{
    "model": "glm-5",
    "messages": [
        {"role": "user", "content": "Your request here"}
    ],
    "thinking": {"type": "enabled"},
    "max_tokens": 4096,
    "temperature": 1.0
}'
```

### cURL - Streaming Call

```bash
curl -X POST "https://api.z.ai/api/paas/v4/chat/completions" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your-api-key" \
-d '{
    "model": "glm-5",
    "messages": [
        {"role": "user", "content": "Your request here"}
    ],
    "thinking": {"type": "enabled"},
    "stream": true,
    "max_tokens": 4096,
    "temperature": 1.0
}'
```

### Python - Official SDK (zai-sdk)

**Installation:**
```bash
pip install zai-sdk
```

**Basic Call:**
```python
from zai import ZaiClient

client = ZaiClient(api_key="your-api-key")

response = client.chat.completions.create(
    model="glm-5",
    messages=[
        {"role": "user", "content": "Your request here"}
    ],
    thinking={"type": "enabled"},
    max_tokens=4096,
    temperature=1.0,
)

print(response.choices[0].message)
```

**Streaming Call:**
```python
from zai import ZaiClient

client = ZaiClient(api_key="your-api-key")

response = client.chat.completions.create(
    model="glm-5",
    messages=[
        {"role": "user", "content": "Your request here"}
    ],
    thinking={"type": "enabled"},
    stream=True,
    max_tokens=4096,
    temperature=0.6,
)

for chunk in response:
    if chunk.choices[0].delta.reasoning_content:
        print(chunk.choices[0].delta.reasoning_content, end="", flush=True)
    if chunk.choices[0].delta.content:
        print(chunk.choices[0].delta.content, end="", flush=True)
```

### Python - OpenAI-Compatible SDK

**Installation:**
```bash
pip install --upgrade 'openai>=1.0'
```

**Usage:**
```python
from openai import OpenAI

client = OpenAI(
    api_key="your-zai-api-key",
    base_url="https://api.z.ai/api/paas/v4/",
)

completion = client.chat.completions.create(
    model="glm-5",
    messages=[
        {"role": "system", "content": "You are a helpful assistant"},
        {"role": "user", "content": "Your request here"}
    ],
)

print(completion.choices[0].message.content)
```

### Python - Legacy SDK (zhipuai)

**Installation:**
```bash
pip install zhipuai==2.1.5.20250726
```

**Usage:**
```python
from zhipuai import ZhipuAI

client = ZhipuAI(api_key="your-api-key")

response = client.chat.completions.create(
    model="glm-5",
    messages=[
        {"role": "user", "content": "Your request here"}
    ],
    thinking={"type": "enabled"},
    max_tokens=65536,
    temperature=1.0
)

print(response.choices[0].message)
```

### Java SDK

**Maven:**
```xml
<dependency>
    <groupId>ai.z.openapi</groupId>
    <artifactId>zai-sdk</artifactId>
    <version>0.3.0</version>
</dependency>
```

**Gradle:**
```groovy
implementation 'ai.z.openapi:zai-sdk:0.3.0'
```

**Usage:**
```java
import ai.z.openapi.ZaiClient;
import ai.z.openapi.service.model.*;

ZaiClient client = ZaiClient.builder()
    .ofZAI()
    .apiKey("your-api-key")
    .build();

ChatCompletionCreateParams request = ChatCompletionCreateParams.builder()
    .model("glm-5")
    .messages(Arrays.asList(
        ChatMessage.builder()
            .role(ChatMessageRole.USER.value())
            .content("Your request here")
            .build()
    ))
    .thinking(ChatThinking.builder().type("enabled").build())
    .maxTokens(4096)
    .temperature(1.0f)
    .build();

ChatCompletionResponse response = client.chat().createChatCompletion(request);
```

---

## Parameter Values from Official Examples

The following values are shown in official documentation examples:

| Parameter | Value in Examples |
|-----------|-------------------|
| `temperature` | `1.0`, `0.6` |
| `max_tokens` | `4096`, `65536` |
| `thinking.type` | `"enabled"`, `"disabled"` |

> **Note:** Z.ai has not published specific parameter recommendations for different use cases. The values above are those shown in official code examples.

---

## Integration with Coding Agents

### Claude Code

Update `~/.claude/settings.json`:

```json
{
  "model": "glm-5",
  "api_base": "https://api.z.ai/api/paas/v4/",
  "api_key": "your-api-key"
}
```

### OpenCode

Create or update `opencode.json`:

```json
{
  "providers": {
    "zai": {
      "api_key": "your-api-key",
      "base_url": "https://api.z.ai/api/paas/v4/"
    }
  },
  "model": "glm-5"
}
```

---

## Supported Languages

- English
- Chinese (中文)

---

## Resources

- **Official Documentation (International):** https://docs.z.ai/guides/llm/glm-5
- **Official Documentation (China):** https://docs.bigmodel.cn/cn/guide/models/text/glm-5
- **API Platform:** https://z.ai/model-api
- **Hugging Face:** https://huggingface.co/zai-org/GLM-5

---

*Last updated: February 2026*
*Source: Z.ai Official Documentation*
