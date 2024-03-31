const args = require('minimist')(process.argv.slice(2))
const QRCode = require("qrcode");
async function main(str) {
    console.log(
        await QRCode.toString(str, { type: "terminal", small: true })
    );
}
main(args['url']);