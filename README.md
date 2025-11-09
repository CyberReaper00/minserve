# MinServe
> [!warning]
> MinServe is created for the purpose of testing and fast development, it is not meant to be used for production or secure environments but for testing, prototyping and hobby projects
>
> Any work being done in those settings should use something like NGINX or Apache Server
---
MinServe is a simple but powerful web server that can be used for the hosting of static pages in a simple and straight-forward manner

## Getting Started
If you wish to try minserve then you can just do the following to build it on your system or go to the [releases page](https://github.com/CyberReaper00/minserve/releases) to download a binary

MinServe is built to be as simple and efficient as possible, due to this simplicity it does not require any `build.sh`, `build.bat`, `run.sh` or anything alike

**Prerequisites**
- Download the main `server.go` file
- Download the Go programming language
- Setup the main environment by doing the following

```
mkdir minserve_proj && cd minserve_proj
go mod init
go get https://github.com/CyberReaper00/helper_utils/humain@latest // this is the main dependency
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
MinServe takes in two arguments,
- Port number
- Filename(s)

The port number is mandatory for obvious reasons, but the filename(s) are optional

If no filename is given then minserve will look in the current folder for the index.html file and set that as the homepage while it also detects and subsequently creates pages for all other html files found in the directory

If any filenames are given then minserve will accept the first name to be the homepage and the rest will be created as top level pages on the site

# Features
- MinServe does provide basic hot-reloading functionality for faster development and can be used in the following manner,
    - Launch minserve with pages
    - Go to the browser and open a page
    - Change the content on that page, locally
    - Save the file and reload the page in the browser
    - The changes will take affect immediately

- MinServe also provides a simple way of shutting down the server by pressing `ctrl-c` in the terminal where minserve is running

- MinServe also has the capability to create sections on the provided port by giving the following syntax,
    - `./new/file.html`

    If the above mentioned syntax is provided to minserve then it will take that to be the homepage like this
    `localhost:[port]/` = `./new/file.html`

    But if the following syntax of `./new/*` is provided then minserve will create the following url(s)
    ```
    localhost:[port]/new/file1
    localhost:[port]/new/file2
    localhost:[port]/new/file3
    localhost:[port]/new/file4
    ```

    This secional syntax can be stacked and an arbitrary amount of sections of the site can be provided to minserve for immediate traversal

    But it must be kept in mind that the first file argument given will always be the homepage

## Not Supported
It does not have support for https requests or writes, so all work - as said before even if it is being done in a professional environment - should only be used for testing new features on the site or prototyping

MinServe also does not support high-traffic throughput, if your site gets thousands of visitors everyday, minserve will not be able to handle such a load and it will most probably crash
