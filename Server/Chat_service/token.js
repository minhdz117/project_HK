var jwt = require ('jsonwebtoken')

const token = jwt.sign({
    "name":"manh"
},"minhdz117", { expiresIn: 24*60*60 });

console.log(token)