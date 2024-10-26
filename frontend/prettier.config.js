/** @type {import("prettier").Config} */
export default {
  plugins: ['@trivago/prettier-plugin-sort-imports', 'prettier-plugin-tailwindcss'],
  semi: false,
  singleQuote: true,
  trailingComma: 'all',
  proseWrap: 'never',
  printWidth: 120,
  tabWidth: 2,
  endOfLine: 'lf',
  importOrderSeparation: true,
  importOrderSortSpecifiers: true,
  importOrder: ['^react$', '<THIRD_PARTY_MODULES>', '^@/(.*)$', '^@test/(.*)$', '^[./]'],
  tailwindConfig: './tailwind.config.ts',
}
