- [Dependency Management](#dependency-management)
  - [GO Path](#go-path)
  - [GO Vendor](#go-vendor)
  - [GO Module](#go-module)


# Dependency Management

## GO Path

- bin : compiled binaries
- pkg : compiled intermediate products to speed up compilation
- src : source code

Disadvantage : Unable to achieve Version Control.

## GO Vendor

- vendor : place a copy of all dependent packages.

Disadvantage : Dependencies conflict.

## GO Module

- go.mod : identify module path and version information (${MAJOR}.{MINOR}.${PATCH}), describe unit dependencies (including labeling indrect and incompatible dependencies). When compiling, go will choose the lowest compatible version.
- Proxy : cache version content to achieve reliable dependency distribution.
- go get/mod : local tools