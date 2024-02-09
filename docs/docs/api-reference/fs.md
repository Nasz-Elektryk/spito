---
sidebar_position: 3
---

# api.fs

The `api.fs` module provides functions for working with the file system.

## pathExists

### Arguments:

- `path` (string): The path to check.

### Returns:

- `exists` (bool): Whether the path exists.

### Example usage:

```lua
exists = api.fs.pathExists("/etc/passwd")
```

## fileExists

### Arguments:

- `path` (string): The path to check.
- `isDirectory` (bool): Whether the path is a directory.

### Returns:

- `exists` (bool): Whether the file exists.

### Example usage:

```lua
exists = api.fs.fileExists("/etc/passwd", false)
```

## readFile

### Arguments:

- `path` (string): The path to read.

### Returns:

- `content` (string): The content of the file.
- `error` (error): The error message if the file does not exist.

### Example usage:

```lua
function readPasswd()
    content, err = api.fs.readFile("/etc/passwd")
    if err ~= nil then
        api.info.error("Error occured during reading the file: " .. err)
        return false
    end
    return true
end
```

## readDir

### Arguments:

- `path` (string): The path to read.

### Returns:

- `files` ([]string): The files in the directory.
- `error` (error): The error message if the directory does not exist.

### Example usage:

```lua
function readDir()
    files = api.fs.readDir("/etc")
    for _, file in ipairs(files) do
        api.info.info(file)
    end
end
```

## fileContains

### Arguments:

- `fileContent` (string): The content of the file.
- `content` (string): The content to check.

### Returns:

- `contains` (bool): Whether the file contains the content.

### Example usage:

```lua
contains = api.fs.fileContains(api.fs.readFile("/etc/passwd"), "root")
```

## removeComments

### Arguments:

- `content` (string): The content to remove comments from.
- `singleLineStart` (string): The start of a single line comment.
- `multiLineStart` (string): The start of a multi line comment.
- `multiLineEnd` (string): The end of a multi line comment.

### Returns:

- `content` (string): The content without comments.

### Example usage:

```lua
content = api.fs.removeComments(api.fs.readFile("/etc/passwd"), "#", "/*", "*/")
```

## find

### Arguments:

- `regex` (string): The regex to search for.
- `fileContent` (string): The content to search in.

### Returns:

- `lines` ([]int): The lines where the regex was found.
- `error` (error): The error message if the regex is invalid.

### Example usage:

```lua
function findRoot()
    lines, err = api.fs.find("root", api.fs.readFile("/etc/passwd"))
    if err ~= nil then
        api.info.error("Error occured during finding the regex: " .. err)
        return false
    end
    return true
end
```

## findAll

### Arguments:

- `regex` (string): The regex to search for.
- `fileContent` (string): The content to search in.

### Returns:

- `lines` ([][]int): The lines where the regex was found.
- `error` - The error message if the regex is invalid.

### Example usage:

```lua
function findAllRoots()
    lines, err = api.fs.findAll("root", api.fs.readFile("/etc/passwd"))
    if err ~= nil then
        api.info.error("Error occured during finding the regex: " .. err)
        return false
    end
    return true
end
```

## getProperLines

### Arguments:

- `regex` (string): The regex to search for.
- `fileContent` (string): The content to search in.

### Returns:

- `lines` ([]string): The lines where the regex was found.
- `error` (error): The error message if the regex is invalid.

### Example usage:

```lua
function getRoots()
    lines, err = api.fs.getProperLines("root", api.fs.readFile("/etc/passwd"))
    if err ~= nil then
        api.info.error("Error occured during finding the regex: " .. err)
        return false
    end
    return true
end
```

## createFile

:::warning
Every file that can be created using this [createConfig](#createconfig) or [updateConfig](#updateconfig)
**should not be created using [createFile](#createfile)**
:::

### Arguments:

- `path` (string): The path to create.
- `content` (string): The content of the file.
- `options` (CreateFileOptions): The options for creating the file.

### Returns:

- `error` (error): The error message if the file already exists.

`CreateFileOptions`:

- `optional` (bool): Whether the file is optional.
- `fileType` (string): The type of the file.

### Example usage:

```lua
function createFile()
    err = api.fs.createFile("/etc/passwd", "root:x:0:0:root:/root:/bin/bash", { optional = false, fileType = "passwd" })
    if err ~= nil then
        api.info.error("Error occured during creating the file: " .. err)
        return false
    end
    return true
end
```

## createConfig

Creates new configuration file or **updates** existing one created using this function.

### Arguments:

- `path` (string): The path to create.
- `content` (string): The content of the file.
- `options` ([CreateConfigOptions](#createconfigoptions)): The options for creating the file.

### Returns:

- `error` (error): The error message.

### Example usage:

```lua
api.fs.createConfig("./eslint.json", '{"root": "false"}', { ConfigType = api.fs.config.json })
```

eslint.json file should look like this:

```json
{
  "root": "false"
}
```

## updateConfig

Similar behaviour to [createConfig](#createconfig), but instead of overriding original
file merges it with the new one.

### Arguments:

- `path` (string): The path to create.
- `content` (string): The content of the file.
- `options` ([CreateConfigOptions](#createconfigoptions)): The options for creating the config.

### Returns:

- `error` (error): The error message.

### Example usage:

```lua
api.fs.updateConfig("./eslint.json", '{"root": "false"}', { ConfigType = api.fs.config.json })
```

eslint.json file should look like this:

```json
{
  "root": "false",
  "other-key": "other-value"
}
```

## compareConfigs

Deeply compares JSON, TOML or other types of [data](#configtype-possible-values-apifsconfig).

### Arguments:

- `received` (string): The content of one file
- `desired` (string): The content of another file
- `configType` ([ConfigType](#configtype-possible-values-apifsconfig)): Format of files

### Returns:

- `error` (error): If nil - they are equal, otherwise they are not or there is a problem

### Example usage:

```lua
api.fs.compareConfigs("./eslint.json", '{"root": "false"}', { ConfigType = api.fs.config.json })
```

## CreateConfigOptions

| Field      | Type       | Description                 |
|------------|------------|-----------------------------|
| Optional   | bool       | The name of the package.    |
| Options    | string     | The version of the package. |
| ConfigType | ConfigType | The version of the package. |

### ConfigType possible values `api.fs.config.*`

- json
- toml
- yaml

### Example usage

```lua
options = {
    Optional = false,
    ConfigType = api.fs.config.toml
}
```