import { swiss } from 'mdx-deck/themes'
import solarized from 'react-syntax-highlighter/styles/prism/solarizedlight'
import prismRust from 'react-syntax-highlighter/languages/prism/rust'

export default {
  ...swiss,
  prism: {
    style: solarized,
    languages: {
      rust: prismRust
    }
  }
}
