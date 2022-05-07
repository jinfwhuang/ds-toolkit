require("@rushstack/eslint-patch/modern-module-resolution");

module.exports = {
    plugins: ["eslint-plugin-tsdoc", "eslint-plugin-eslint-comments", "mocha"],
    extends: ["eslint:recommended", "prettier"],
    ignorePatterns: ["**/dist/*", "**/dist-cjs/*", "**/coverage/*"],
    parserOptions: {
        ecmaVersion: 8,
        sourceType: "module"
    },
    env: {
        node: true,
        es6: true
    },
    rules: {
        camelcase: "error",
        "one-var-declaration-per-line": ["error", "always"],
        "no-implicit-coercion": [
            2,
            {
                boolean: true
            }
        ],
        "object-shorthand": ["error", "always"],
        "prefer-spread": "error",
        "prefer-arrow-callback": "error",
        "new-cap": [
            "error",
            {
                newIsCap: true
            }
        ],
        "comma-dangle": ["error", "never"],
        "mocha/no-exclusive-tests": "error",
        "eslint-comments/require-description": ["warn", { ignore: [] }],
        "prefer-promise-reject-errors": "error",
        "eqeqeq": "error",
        "no-throw-literal": "error"
    }
};
