## What is Viper?

Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats. It supports:


<details><summary>show</summary>
<p>

```
* setting defaults
* reading from JSON, TOML, YAML, HCL, and Java properties config files
* live watching and re-reading of config files (optional)
* reading from environment variables
* reading from remote config systems (etcd or Consul), and watching changes
* reading from command line flags
* reading from buffer
* setting explicit values

Viper can be thought of as a registry for all of your applications configuration needs.

```

</p>
</details>

---

## Why Viper?

When building a modern application, you don’t want to worry about configuration file formats; you want to focus on building awesome software. Viper is here to help with that.

Viper does the following for you:


<details><summary>show</summary>
<p>

```
* Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, or Java properties formats.
* Provide a mechanism to set default values for your different configuration options.
* Provide a mechanism to set override values for options specified through command line flags.
* Provide an alias system to easily rename parameters without breaking existing code.
* Make it easy to tell the difference between when a user has provided a command line or config file which is the same as the default.


```

</p>
</details>

---

## Viper uses the following precedence order. Each item takes precedence over the item below it:


<details><summary>show</summary>
<p>

```
* explicit call to Set
* flag
* env
* config
* key/value store
* default

Viper configuration keys are case insensitive.


```

</p>
</details>

---

## Establishing Defaults

A good configuration system will support default values. A default value is not required for a key, but it’s useful in the event that a key hasn’t been set via config file, environment variable, remote configuration or flag.

Examples:

<details><summary>show</summary>
<p>

```
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

```

</p>
</details>

---

## Watching and re-reading config files

Viper supports the ability to have your application live read a config file while running.

Make sure you add all of the configPaths prior to calling WatchConfig()

<details><summary>show</summary>
<p>

```
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)
})

```

</p>
</details>

---

## Env example

<details><summary>show</summary>
<p>

```
SetEnvPrefix("spf") // will be uppercased automatically
BindEnv("id")

os.Setenv("SPF_ID", "13") // typically done outside of the app

id := Get("id") // 13

```

</p>
</details>

---

## Viper- Individual flag binding

For individual flags, the BindPFlag() method provides this functionality.

Example:

<details><summary>show</summary>
<p>

```
serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))


```

</p>
</details>

---

## Viper- Binding to existing flags

You can also bind an existing set of pflags (pflag.FlagSet):

Example:


<details><summary>show</summary>
<p>

```
pflag.Int("flagname", 1234, "help message for flagname")

pflag.Parse()
viper.BindPFlags(pflag.CommandLine)

i := viper.GetInt("flagname") // retrieve values from viper instead of pflag


```

</p>
</details>

---

## Remote Key/Value Store Support

To enable remote support in Viper, do a blank import of the viper/remote package:

import _ "github.com/spf13/viper/remote"

<details><summary>show</summary>
<p>

```
You can use remote configuration in conjunction with local configuration, or independently of it.

crypt has a command-line helper that you can use to put configurations in your K/V store. crypt defaults to etcd on http://127.0.0.1:4001.

```

</p>
</details>

---

## Title

<details><summary>show</summary>
<p>

```


```

</p>
</details>

---

## Title

<details><summary>show</summary>
<p>

```
go get github.com/xordataexchange/crypt/bin/crypt
crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json

# Confirm that your value was set:

crypt get -plaintext /config/hugo.json

```

</p>
</details>

---

## Remote Key/Value Store Example - Unencrypted

Etcd

<details><summary>show</summary>
<p>

```
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"
err := viper.ReadRemoteConfig()

```

</p>
</details>

---

## Remote Key/Value Store Example - Encrypted

<details><summary>show</summary>
<p>

```
viper.AddSecureRemoteProvider("etcd","http://127.0.0.1:4001","/config/hugo.json","/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // because there is no file extension in a stream of bytes,  supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"
err := viper.ReadRemoteConfig()

```

</p>
</details>

---

## Watching Changes in etcd - Unencrypted

<details><summary>show</summary>
<p>

```
// alternatively, you can create a new viper instance.
var runtime_viper = viper.New()

runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
runtime_viper.SetConfigType("yaml") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"

// read from remote config the first time.
err := runtime_viper.ReadRemoteConfig()

// unmarshal config
runtime_viper.Unmarshal(&runtime_conf)

// open a goroutine to watch remote changes forever
go func(){
	for {
	    time.Sleep(time.Second * 5) // delay after each request

	    // currently, only tested with etcd support
	    err := runtime_viper.WatchRemoteConfig()
	    if err != nil {
	        log.Errorf("unable to read remote config: %v", err)
	        continue
	    }

	    // unmarshal new config into our runtime config struct. you can also use channel
	    // to implement a signal to notify the system of the changes
	    runtime_viper.Unmarshal(&runtime_conf)
	}
}()

```

</p>
</details>

---

## Getting Values From Viper

In Viper, there are a few ways to get a value depending on the value’s type. The following functions and methods exist:

<details><summary>show</summary>
<p>


* Get(key string) : interface{}
* GetBool(key string) : bool
* GetFloat64(key string) : float64
* GetInt(key string) : int
* GetString(key string) : string
* GetStringMap(key string) : map[string]interface{}
* GetStringMapString(key string) : map[string]string
* GetStringSlice(key string) : []string
* GetTime(key string) : time.Time
* GetDuration(key string) : time.Duration
IsSet(key string) : bool
AllSettings() : map[string]interface{}



</p>
</details>

---

## Accessing nested keys

The accessor methods also accept formatted paths to deeply nested keys. For example, if the following JSON file is loaded:



<details><summary>show</summary>
<p>


```
{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}



```

Viper can access a nested field by passing a . delimited path of keys:

GetString("datastore.metric.host") // (returns "127.0.0.1")

</p>
</details>


