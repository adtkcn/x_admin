import js from '@eslint/js'
import globals from 'globals'
import tseslint from 'typescript-eslint'
import pluginVue from 'eslint-plugin-vue'
import json from '@eslint/json'
import css from '@eslint/css'

export default [
    tseslint.configs.recommended,
    pluginVue.configs['flat/essential'],

    {
        files: ['**/*.{js,mjs,cjs,ts,mts,cts,vue}'],
        plugins: { js },
        extends: ['js/recommended'],
        languageOptions: { globals: globals.browser },
        rules: {
            'no-undef': 'off',
            '@typescript-eslint/no-unused-expressions': 'off',
            '@typescript-eslint/ban-ts-comment': 'off',
            'no-unused-vars': 'off'
            // '@typescript-eslint/no-unused-vars': 'off'
        }
    },

    {
        files: ['**/*.vue'],
        languageOptions: { parserOptions: { parser: tseslint.parser } },
        rules: {
            'vue/multi-word-component-names': 'off',
            '@typescript-eslint/no-explicit-any': 'off',
            'no-empty': 'off',
            'no-undef': 'off',
            '@typescript-eslint/no-unused-vars': 'off'
        }
    },
    {
        files: ['**/*.json'],
        plugins: { json },
        language: 'json/json',
        extends: ['json/recommended']
    },
    {
        files: ['**/*.jsonc'],
        plugins: { json },
        language: 'json/jsonc',
        extends: ['json/recommended']
    },
    {
        files: ['**/*.json5'],
        plugins: { json },
        language: 'json/json5',
        extends: ['json/recommended']
    },
    { files: ['**/*.css'], plugins: { css }, language: 'css/css', extends: ['css/recommended'] }
]
