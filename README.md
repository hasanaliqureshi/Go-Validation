![Image of Yaktocat](https://www.devteam.space/wp-content/uploads/2017/03/gopher_head-min.png
)

# Text and Input Validation in Golang

I've created this helping script for my APIs to deal with user inputs. Just clone the repo and use it in your project.

## Usage:

1. Define the validator

  `
  userInput := validator("blah!blah!blah!")
  `

2. Apply the methods

   `newInput.Required().oneLowerCase().check()`     `// will return true or false` 



## Methods

use `.check()` after all methods to return the `boolean` value, otherwise the method will return `type struct`

| Methods            | Descrption                                                   |
| ------------------ | :----------------------------------------------------------- |
| Required()         | returns `false` if input is empty                            |
| minLength(int)     | returns `false` if input is less than provided minimum value |
| maxLength(int)     | returns `false` if input is greater than provided maximum value |
| isEmail            | return `false` if input is not an email address              |
| oneLowerCase       | return `false` if input value doesn't have at least one lowercase character |
| allLowerCase       | return `false`if input value doesn't have all lowercase characters |
| oneUpperCase       | return `false` if input value doesn't have at least one uppercase character |
| allUpperCase       | return `false` if input value doesn't have all uppercase characters |
| oneNumber          | return `false` if input value doesn't have at least one number character |
| isSpecialCharacter | return `false` if input value doesn't have at least one special character |


