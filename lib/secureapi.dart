import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:Auth.APP/pathlookup.dart';

Future<HttpRequest> sendLogin(
    Map<String, String> app, String email, String password) async {
  final url = await buildPath("Secure.API", "login", new List<String>());
  final data = jsonEncode({"App": app, "Email": email, "Password": password});

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

Future<HttpRequest> sendForgot(String identity) async {
  final url = await buildPath("Secure.API", "forgot", new List<String>());
  final data = jsonEncode(identity);

  final compltr = new Completer<HttpRequest>();
  final request = HttpRequest();
  request.open("POST", url);
  request.setRequestHeader("Content-Type", "text/json;charset=UTF-8");
  request.onLoadEnd
      .listen((e) => compltr.complete(request), onError: compltr.completeError);
  request.onError.listen(compltr.completeError);
  request.onProgress.listen(onProgress);
  request.send(data);

  return compltr.future;
}

void onProgress(ProgressEvent e) {
  if (e.lengthComputable) {
    print('Progress... ${e.total}/${e.loaded}');
  }
}
