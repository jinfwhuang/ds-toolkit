
## Overview

This is an overview of the global tsconfig

## Compiler options

A brief overview of each compiler option is below

- incremental - caches info about prev compiled project graph to speed up future builds
- composite - speeds up build process
- declarationMap - generates a map between declaration files and the original files. Helps vscode display things
- moduleResolution - specifies how the compiler finds and resolves module imports - basically defines how module resolution works for inputs for the compiler (how you should write imports/exports)
- module - specifies how the module resolution will work in the output of the compiler
- target - specifies what version of javascript tsc will compile the typescript into
- sourceMap - generates a map file allowing debuggers to display original javascript when actually working with the compiled version
- allowJs - disallows js files
- noEmit - certain types of projects bug out if noEmit is not declaratively defined, easier to just have it defined in root
- strict - enforces strict type checking by default
- forceConsistentCasingInFileNames
- jsx - This determines how typescript handles jsx. It leaves it for metro transform step to handle.
- isolatedModules - Requires all files to be modules (must have an import/export). Non-modules pollute global state for the program.
- esModuleInterop - fixes a couple errors that can occur during default TS module resolution
- noImplicitReturns - requires explicit return types - not specified as defaulted to true
- noUnusedLocals - enforces good practice
- noUnusedParams - enforces good practice
- resolveJsonModule - TS does not resolve json by default

For more information, go to the documentation page: https://www.typescriptlang.org/tsconfig