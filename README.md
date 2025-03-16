![Banner](/rsc/banner.png)
# A.S.K. – AI Script Kiddie (DeepSeek + Powershell)

This project is a simple command-line application that queries the DeepSeek API to find the best PowerShell or Azure CLI cmdlet for a specific task.
 It interacts with the DeepSeek service to provide the best answer straight into your terminal, so you don't need to keep going back and forth within your browser.
 This is a quick project for the sole purpose of practicing compiling Go projects and also experimenting with the DeepSeek API. 
 The application is ideal for DevOps professionals,users that are learning Powershell, and anyone working with Windows OS and related technologies.

## Table of Contents
- Overview
- Installation
- Usage
- Environment Variables
- Dependencies


## Overview
The DeepSeek PowerShell Command Finder is built in Go and uses the DeepSeek API to find and return the most relevant cmdlets for user input. The application asks the user for a task-related question, sends the query to the DeepSeek API, and then outputs the appropriate cmdlet based on the best response.

## Features:
Queries DeepSeek API for cmdlets related to PowerShell, Azure CLI, and DevOps.
Outputs only the cmdlet (no extra information).
Helps users find the best cmdlet for a given task in Windows OS, PowerShell, or DevOps.
Installation
To run this project, you need Go installed on your local machine.

![screenShot1](/rsc/ss2.png)
![screenShot1](/rsc/ss1.png)
## Build yourself:
```
git clone https://github.com/your-username/deepseek-cmdlet-finder.git
cd ask
go build -o ask.exe
```

## Pre-compiled
Simply add the **ask.exe** into your user paths and you can run it from the terminal by typing:
```
ask
```

## Usage
The application will prompt you to enter a task or question related to PowerShell, Azure CLI, or DevOps.
Enter your query in the command-line interface (CLI).
The application will return the best cmdlet that corresponds to your query, based on DeepSeek's AI-powered response.

Example:
```
» I need to create a VM on azure that is isolated from the internet and has ISS service and MySQL pre-installed
```
Output:
```
New-AzVm -ResourceGroupName "MyResourceGroup" -Name "MyVM" -Image "MicrosoftWindowsServer:WindowsServer:2019-Datacenter:latest" -VirtualNetworkName "MyVnet" -SubnetName "MySubnet" -PublicIpAddressName "MyPublicIp" -SecurityGroupName "MyNsg" -OpenPorts 80,3306 -AsJob -CustomData @' ##cloud-config package_update: true packages: - iis - mysql-server runcmd: - systemctl enable mysqld - systemctl start mysqld '@
```
## Environment Variables

For the application to work you will need to have an DeepSeek API on your system variable called **apiDeepSeek**
Please note you will also need to have some money on your account, API calls are not free lol, but $2.00 is enough to last you a long time.

## Dependencies
The project uses the following Go packages:

github.com/go-deepseek/deepseek: A Go client for interacting with the DeepSeek API.
github.com/go-deepseek/deepseek/request: A Go package for making requests to the DeepSeek API.
These dependencies are automatically included when you run go build.

rsrc_windows_amd64.syso is the icon file compiled as a Windows Resource
```
go install github.com/akavel/rsrc@latest
rsrc -ico icon.ico
```
