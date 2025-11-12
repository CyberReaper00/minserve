# MinServe
> [!warning]
> MinServe is created for the purpose of testing and fast development, it is not meant to be used for production or secure environments but for testing, prototyping and hobby projects
> It must also be known that MinServe is heavily under construction, is not complete in any way and will include a lot of problems
>
> Any work being done in those settings should use something like NGINX or Apache Server
---
MinServe is a simple but powerful web server that can be used for the hosting of static pages in a simple and straight-forward manner

## Getting Started
If you wish to try minserve then you can just do the following to build it on your system or go to the [releases page](https://github.com/CyberReaper00/minserve/releases) to download a binary

All release binaries are statically built because that is the correct way to build software, if you want the binaries to be dynamic then you can build them yourself with the following guide

MinServe is built to be as simple and efficient as possible, due to this simplicity it does not require any `build.sh`, `build.bat`, `run.sh` or anything alike

**Prerequisites**
- Download the Go programming language
- Clone the repo or download the source code
- Setup the main environment by doing the following

```
mkdir minserve_proj && cd minserve_proj
go mod init
go mod tidy
```

**Building**
- Once everything has been set up, run the following command for your OS

### Linux
**Dynamic linking:**
```bash
go build -o minserve .
```

**Static linking:**
```bash
CGO_ENABLED=0
go build -o minserve .
```

### Windows or MacOS
All go binaries on windows and macos are as statically linked by default as they allow them to be, there really isnt anything else to be done
```bash
go build -o minserve .
```

## Usage
MinServe takes in the port to host at
`minserve 1234`

It can be used in the following manner,
- Launch minserve with port
- Go to the browser and open a page
- Change the content on that page, locally
- Save the file and reload the page in the browser
- The changes will take affect on reload

Currently minserve hosts all files in the current directory and all subdirectories, it does not allow for any exclusions
Any file in the current directory is displayed as such,
```
localhost:1234/file1
localhost:1234/file2
```

Any file in any subdirecory is displayed as such,
```
localhost:1234/folder1/file1
localhost:1234/folder1/file2

localhost:1234/folder2/file1
localhost:1234/folder2/file2
```

# Features
- The main reason for MinServe to exist is to provide basic hot-reloading functionality that the default Go server does not provide
    - More functionality will be added with time as seen fit
- If an `index.html` file is not found then it will give an error and exit
- It includes an implementation for a custom 404 page which can be updated and customized at will and an example file has been provided
`page_not_found`
- It looks for the word `index` in the directory for the homepage, if you have any other files named index like `index.js` then minserve will get confused and the webpage will not be able to respond to the servers request

Anything that is not mentioned is not provided or supported by minserve
