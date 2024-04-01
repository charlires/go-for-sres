# The good of Go for SREs

If you are an SRE, Go must be one of the languages you should know for sure.
There are multiple ways you can use Go in your day to day.

In this repository you will find few reasons why Go is one of the best languages you must master as an SRE

## CLI tools

### Standard Library

In Go is simple to create you own CLI tool. 

Using the standard library you can create a basic CLI 
[EXAMPLE](./simple-cli/main.go) 
```sh 
cd simple-cli
go build .
./simple-cli --name carlos
```

### Cobra 
https://github.com/spf13/cobra   

Using Cobra you can create complex command line tools with many features already build in

Cobra is used in [helm](https://github.com/helm/helm/blob/main/cmd/helm/root.go), hugo, gh

With Cobra you get these features out of the box:
- Help output
- Autocompletion for zsh or fish
- command suggestion "Do you mean this?"
- Command flags	
- Multi platform binaries

Why not use bash scripts, it's not intuitive, personally I always struggle reading bash files, there is no standard way to code it, there is no intellisense, there is no easy way to test them.

[EXAMPLE](./cobra-cli/main.go)
```sh 
cd cobra-cli
go build .
./cobra-cli echo -t=4 carlos
```

## File Templating

### Go Standard File Templates
Go has a powerfull templating engine 
[EXAMPLE](./templates/main.go)

### HELM Templates
This template engine has been used by other tools like[helm](./helm/mychart/templates/configmap.yaml)   
HELM takes full advantage of the Go template engine  
<https://www.practical-go-lessons.com/chap-32-templates>  
<https://helm.sh/docs/chart_template_guide/named_templates>

[EXAMPLE](./helm/mychart/templates/configmap.yaml)
```sh
helm template --dry-run ./mychart/
```

## Web server
Go is one of the languages where creating an standalone web server is simple and secure.
[EXAMPLE](./web-server/main.go)

## Infrastructure as Code

With Go you can create you own infrastructure using Pulumi 

<https://github.com/hashicorp/terraform-plugin-go>   
<https://github.com/hashicorp/terraform-exec>   
<https://developer.hashicorp.com/terraform/cdktf/api-reference/go>  

### unittest infrastructure as code
You can test you Terraform plans using Go with [Terratest](./terratest/main.tf)

## Honorable mentions
https://docs.dagger.io/manuals/developer/go 


## References
https://roadmap.sh/golang   
https://github.com/avelino/awesome-go   