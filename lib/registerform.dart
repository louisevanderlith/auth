import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_entity/bodies/register.dart';
import 'package:mango_entity/secureapi.dart';
import 'package:mango_ui/formstate.dart';

class RegisterForm extends FormState {
  TextInputElement _name;
  EmailInputElement _email;
  PasswordInputElement _password;
  PasswordInputElement _confirm;

  RegisterForm(String idElem, String nameElem, String emailElem,
      String passElem, String confirmElem, String submitBtn)
      : super(idElem, submitBtn) {
    _name = querySelector(nameElem);
    _email = querySelector(emailElem);
    _password = querySelector(passElem);
    _confirm = querySelector(confirmElem);

    querySelector(submitBtn).onClick.listen(onSend);
  }

  String get name {
    return _name.value;
  }

  String get email {
    return _email.value;
  }

  String get password {
    return _password.value;
  }

  String get confirmPassword {
    return _confirm.value;
  }

  void onSend(Event e) {
    if (isFormValid() && passwordsMatch()) {
      disableSubmit(true);
      submitSend().then((obj) {
        disableSubmit(false);
      });
    }
  }

  Future submitSend() async {
    var data = new Register(name, email, password, confirmPassword);
    var req = await sendRegister(data);

    var obj = jsonDecode(req.response);

    print(obj['Data']);
    afterSend(obj['Data']);

    if (req.status == 200) {
      new Toast.success(
          title: "Success!",
          message: req.response,
          position: ToastPos.bottomLeft);
      super.form.reset();
    } else {
      new Toast.error(
          title: "Error!",
          message: req.response,
          position: ToastPos.bottomLeft);
    }
  }

  bool passwordsMatch() {
    return password == confirmPassword;
  }

  void afterSend(Object obj) {
    print("We have touchdown {obj}");
  }
}
