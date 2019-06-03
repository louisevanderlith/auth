import 'dart:convert';
import 'dart:html';
import 'package:Auth.APP/secureapi.dart';

import 'formstate.dart';
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

  void onSend(Event e) async {
    if (isFormValid()) {
      disableSubmit(true);

      final app = await getApp();
      var result = await sendLogin(app, email, password);
      var obj = jsonDecode(result.response);

      if (result.status == 200) {
        afterSend(obj['Data']);
      } else {
        _error.text = obj['Error'];
      }
    }
  }
  
  void afterSend(String sessionID) {
    var finalURL = window.localStorage['return'];
    finalURL += "?access_token=" + sessionID;

    window.location.replace(finalURL);
  }
}
