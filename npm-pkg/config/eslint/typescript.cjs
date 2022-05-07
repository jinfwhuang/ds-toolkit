module.exports = {
    parser: "@typescript-eslint/parser",
    plugins: ["@typescript-eslint", "import", "monorepo-cop"],
    extends: [
        "./javascript.cjs",
        "plugin:@typescript-eslint/recommended",
        "prettier/@typescript-eslint",
        "plugin:import/typescript",
        "plugin:monorepo-cop/recommended"
    ],
    rules: {
        "@typescript-eslint/array-type": "error",
        "@typescript-eslint/ban-ts-comment": "off",
        "@typescript-eslint/explicit-module-boundary-types": "off",
        "@typescript-eslint/no-empty-interface": ["error", { allowSingleExtends: true }],
        "@typescript-eslint/no-explicit-any": "error",
        "@typescript-eslint/no-non-null-assertion": "off",
        "@typescript-eslint/no-unused-vars": ["error", { argsIgnorePattern: "^_" }],
        "@typescript-eslint/no-unused-vars-experimental": ["error"],
        "@typescript-eslint/naming-convention": [
            "error",
            {
                selector: "parameter",
                format: null,
                modifiers: ["unused"],
                leadingUnderscore: "require"
            },
            {
                selector: "parameter",
                format: null,
                leadingUnderscore: "forbid"
            }
            /*
            // we should add this to enforce private field naming
            {
                selector: "classProperty",
                format: null,
                modifiers: ["private"],
                leadingUnderscore: "require"
            }
            */
        ],
        "tsdoc/syntax": "error",
        "import/no-default-export": "error",
        "import/extensions": [
            "error",
            "ignorePackages",
            {
                js: "always"
            }
        ],
        "no-restricted-imports": [
            "error",
            {
                patterns: ["*/dist/*.js", "*/dist/*"]
            }
        ],
        "no-await-in-loop": "error"
    }
};
