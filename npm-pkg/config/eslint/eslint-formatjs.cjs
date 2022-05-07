module.exports = {
    //https://formatjs.io/docs/tooling/linter#usage
    "plugins": ["formatjs"],
    "rules": {
      "formatjs/enforce-default-message": ["error", "literal"],
      "formatjs/enforce-plural-rules": [
        "error",
        {
          "one": true,
          "other": true,
          "zero": false
        }
      ],
      //@hz/<packagename>:<string-id>
      "formatjs/enforce-id": [
        "error"
      ]
    }
};
