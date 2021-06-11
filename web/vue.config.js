module.exports = {
  configureWebpack: config => {
    config.devServer = {
      headers: {
        'X-Sample-Header': "Sample"
      }
    }
   }
}