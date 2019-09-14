/**
 * Кравлер
 *
 * @module @nitra/crawler
 */

// @ts-ignore
const jsonData = require('./main.json')

/**
 * accept-language preferred value by domain TLD
 *
 * @param {String} domain - для кого
 * @result {String}
 */
function Language (domain) {
  try {
    const parsed = new URL(domain)
    if (parsed.hostname) {
      domain = parsed.hostname
    }
  } catch (err) {
    ;
  }

  if (domain.substr(-3) === '.ua') {
    return jsonData[domain.substr(-2)]
  }
}

// Export it to make it available outside
module.exports = Language
