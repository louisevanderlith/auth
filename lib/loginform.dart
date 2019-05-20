import 'dart:async';
import 'dart:convert';
import 'dart:html';
import 'formstate.dart';
import 'pathlookup.dart';
import 'app.dart';

class LoginForm extends FormState {
  EmailInputElement _email;
  PasswordInputElement _password;
  ParagraphElement _error;

  LoginForm(
      String idElem, String emailElem, String passwordElem, String submitBtn)
      : super(idElem, submitBtn) {
    _email = querySelector(emailElem);
    _password = querySelector(passwordElem);
    _error = querySelector("#frmError");

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get email {
    return _email.value;
  }

  String get password {
    return _password.value;
  }

  void onSend(Event e) {
    if (isFormValid()) {
      disableSubmit(true);
      submitSend().then((x) {
        disableSubmit(false);
      });
    }
  }

  Future submitSend() async {
    var url = await buildPath("Secure.API", "login", new List<String>());
    var data = jsonEncode(
        {"App": await getApp(), "Email": email, "Password": password});

    HttpRequest.request(url,
            method: "POST", sendData: data, onProgress: loginProgress)
        .then((HttpRequest req) {
      var obj = jsonDecode(req.response);

      if (obj['Error'] != "") {
        _error.text = obj['Error'];
        return;
      }

      afterSend(obj['Data']);
    }).catchError((e) {
      _error.text = e.error;
    });
  }

  void loginProgress(ProgressEvent e) {
    if (e.lengthComputable) {
      print('Progress... ${e.total}/${e.loaded}');
    }
  }

  void afterSend(String sessionID) {
    var finalURL = window.localStorage['return'];
    finalURL += "?access_token=" + sessionID;

    window.location.replace(finalURL);
  }
}
