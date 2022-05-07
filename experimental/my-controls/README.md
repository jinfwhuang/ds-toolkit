
## Overview

Sample package that can be used by other apps or packages within this workspace

## Commands
```bash
yarn build
```


## How to create a package

1. Initialize the package


2. Update your package.json

    ```
    "main": "./dist/index.js",
    "typings": "./dist/index.d.ts",
    "scripts": {
        "build": "tsc"
    },
    ```

3. Add a tsconfig.json file with the following content:
   
    ```
    {
        "compilerOptions": {
            "module": "commonjs",
            "rootDir": "<YourRootDirHere>",
            "outDir": "./dist",
            "sourceMap": true,
            "declaration": true,
            "strictPropertyInitialization": false,
            "skipLibCheck": true
        },
    }

    ```

4. Build your package by running this in terminal `yarn build`



### TODO

Configure global yarn workspaces tsconfig to avoid conflicting settings. The initial settings of this tsconfig was not working with the react native app. This should be handled at the monorepo level.