const Language = require('./index')

/* global test expect */
test('Get expiration for nitra.ai', () => {
  expect(Language('http://nitra.ai/')).toBe(undefined)
  expect(Language('nitra.ua')).toBe('uk-UA, uk;q=0.9, ru;q=0.8, *;q=0.5')
})
