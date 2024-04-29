module.exports = {
    env: {
        browser: true,
        es2021: true,
        node: true,
    },
    extends: ['eslint:recommended', 'plugin:solid/recommended', 'plugin:@typescript-eslint/recommended', 'plugin:prettier/recommended'],
    overrides: [],
    parser: '@typescript-eslint/parser',
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
    },
    plugins: ['solid', 'prettier', '@typescript-eslint'],
    rules: {
        '@typescript-eslint/ban-ts-comment': 'off',
        '@typescript-eslint/no-explicit-any': 'off',
        '@typescript-eslint/ban-types': 'off',
        '@typescript-eslint/no-unused-vars': 'off',
        'solid/no-react-specific-props': 'off',
        'solid/jsx-no-undef': 'off',
        'prefer-const': 'off'
    },
};
