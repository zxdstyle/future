import antfu from "@antfu/eslint-config"

export default antfu({
}, {
    rules: {
        "no-alert": "off",
        "no-console": "off",
        "indent": ["error", 4],
        "style/indent": ["error", 4, { SwitchCase: 0 }],
        "style/jsx-indent": ["error", 4],
        "style/jsx-indent-props": ["error", 4],
        "style/jsx-closing-tag-location": "off",
        "node/prefer-global/process": "off",
        "max-len": ["error", 160],
        "quotes": ["error", "double"],
        "style/quotes": ["error", "double"],
        "array-callback-return": "off",
        "ts/ban-ts-comment": "off",
    },
})
