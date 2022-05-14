/**
 * Remove certificate pinning
 * Usage: frida --enable-jit -U -f com.ultimateguitar.tabs -l ./scripts/unpin.js --no-pause
 */
 if (Java.available) {
  console.log('[*] logger.js - Java is available');
  Java.perform(function () {
      var client = Java.use("okhttp3.OkHttpClient");
      var call = Java.use("okhttp3.RealCall");

      client.newCall.implementation = function (request) {
        var result = this.newCall(request);
        console.log('[*] logger.js - ', request.toString());
        return result;
      };

      call.getResponseWithInterceptorChain.implementation = function () {
        var response = this.getResponseWithInterceptorChain();
        console.log('[*] logger.js - response: ', response.toString());
        // Capture response as-needed.
        // console.log('response: ', response.body().string());
        return response;
      };

      var AppUtils = Java.use("com.ultimateguitar.ugpro.utils.AppUtils");
      AppUtils.getApiKey.implementation = function () {
        console.log('[*] logger.js - getApiKey called!', s);
        return this.getApiKey();
      };
      
      AppUtils.getMd5.implementation = function (s) {
        console.log('[*] logger.js - getMd5 called!', s);
        return this.getMd5(s);
      };
  });
}


