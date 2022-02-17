# How to run this file:
#  PS> iwr -useb https://raw.githubusercontent.com/JayBazuzi/machine-setup/main/dev_environments/golang.ps1 | iex

iwr -useb https://raw.githubusercontent.com/JayBazuzi/machine-setup/main/windows.ps1 | iex
iwr -useb https://raw.githubusercontent.com/JayBazuzi/machine-setup/main/golang-goland.ps1 | iex

# Clone repo
& "C:\Program Files\Git\cmd\git.exe" clone https://github.com/OrangePants-R/starter-project-go-coderetreat.git C:\Code\starter-project-go-coderetreat
