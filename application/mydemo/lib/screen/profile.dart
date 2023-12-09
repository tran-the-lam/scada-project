import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';
import 'package:mydemo/utils/http.dart' as http;
import 'package:mydemo/utils/utils.dart' as utils;

enum SensorType { temperature, humidity, pressure }

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key, required this.title});
  final String title;

  @override
  State<ProfilePage> createState() => _ProfileBody();
}

class _ProfileBody extends State<ProfilePage> {
  TextEditingController oldPwd = TextEditingController();
  TextEditingController newPwd = TextEditingController();
  TextEditingController confirmPwd = TextEditingController();

  bool _isOldPwdVisible = false;
  bool _isNewPwdVisible = false;
  bool _isConfirmPwdVisible = false;
  String userID = '';

  @override
  void initState() {
    super.initState();
    fetchUserID();
  }

  Future<void> fetchUserID() async {
    final userID = await utils.getUserId();
    setState(() {
      this.userID = userID;
    });
  }

  Future<void> updatePwdOnPress() async {
    if (oldPwd.text.isEmpty || newPwd.text.isEmpty || confirmPwd.text.isEmpty) {
      Fluttertoast.showToast(
          msg: "Vui lòng nhập đầy đủ thông tin",
          toastLength: Toast.LENGTH_LONG,
          gravity: ToastGravity.BOTTOM,
          timeInSecForIosWeb: 1,
          backgroundColor: Colors.redAccent,
          textColor: Colors.white,
          webPosition: "center",
          fontSize: 16.0);
      return;
    } else if (confirmPwd.text != newPwd.text) {
      Fluttertoast.showToast(
          msg: "Mật khẩu không khớp",
          toastLength: Toast.LENGTH_LONG,
          gravity: ToastGravity.BOTTOM,
          timeInSecForIosWeb: 1,
          backgroundColor: Colors.redAccent,
          textColor: Colors.white,
          webPosition: "center",
          fontSize: 16.0);
      return;
    }

    //TODO call API to change password
    final isSuccess = await http.updatePassword(oldPwd.text, newPwd.text);
    if (isSuccess) {
      Fluttertoast.showToast(
        msg: "Đổi mật khẩu thành công",
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.BOTTOM,
        timeInSecForIosWeb: 1,
        backgroundColor: Colors.greenAccent,
        textColor: Colors.white,
        webPosition: "center",
        fontSize: 16.0,
      );
      Navigator.of(context).pop();
    } else {
      Fluttertoast.showToast(
        msg: "Đổi mật khẩu thất bại",
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.BOTTOM,
        timeInSecForIosWeb: 1,
        backgroundColor: Colors.redAccent,
        textColor: Colors.white,
        webPosition: "center",
        fontSize: 16.0,
      );
    }
  }

  void showAddAcountPopup(BuildContext context) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return StatefulBuilder(
          builder: (context, setState) {
            return AlertDialog(
              title: Text('Thay đổi mật khẩu',
                  style: TextStyle(fontWeight: FontWeight.bold),
                  textAlign: TextAlign.center),
              content: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  TextField(
                    obscureText: !_isOldPwdVisible,
                    controller: oldPwd,
                    decoration: InputDecoration(
                      hintText: 'Nhập mật khẩu cũ',
                      suffixIcon: IconButton(
                      icon: Icon(
                        _isOldPwdVisible ? Icons.visibility : Icons.visibility_off,
                      ),
                      onPressed: () {
                        setState(() {
                          _isOldPwdVisible = !_isOldPwdVisible;
                        });
                      },
                    ),
                  ),
                  ),
                  SizedBox(height: 10, width: 10),
                  TextField(
                    obscureText: !_isNewPwdVisible,
                    controller: newPwd,
                    decoration: InputDecoration(
                      hintText: 'Nhập mật khẩu mới',
                      suffixIcon: IconButton(
                      icon: Icon(
                        _isNewPwdVisible ? Icons.visibility : Icons.visibility_off,
                      ),
                      onPressed: () {
                        setState(() {
                          _isNewPwdVisible = !_isNewPwdVisible;
                        });
                      },
                    ),
                    ),
                  ),
                  SizedBox(height: 10, width: 10),
                  TextField(
                    obscureText: !_isConfirmPwdVisible,
                    controller: confirmPwd,
                    decoration:
                        InputDecoration(hintText: 'Xác nhận lại mật khẩu',
                        suffixIcon: IconButton(
                      icon: Icon(
                        _isConfirmPwdVisible ? Icons.visibility : Icons.visibility_off,
                      ),
                      onPressed: () {
                        setState(() {
                          _isConfirmPwdVisible = !_isConfirmPwdVisible;
                        });
                      },
                    ),
                    ),
                  ),
                  SizedBox(height: 10, width: 10),
                ],
              ),
              actions: [
                ElevatedButton(
                  onPressed: updatePwdOnPress,
                  child: Text('Đồng ý'),
                ),
                ElevatedButton(
                  onPressed: () {
                    // Đóng popup khi nhấn nút Hủy
                    Navigator.of(context).pop();
                  },
                  child: Text('Hủy'),
                ),
              ],
            );
          },
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Profile',
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(
          title: Text('Profile'),
        ),
        body: Center(
          child: Column(
            // mainAxisAlignment: MainAxisAlignment.center,
            children: [
              SizedBox(height: 70),
              CircleAvatar(
                radius: 50,
                backgroundImage: const AssetImage('../assets/avatar.png'),
              ),
              SizedBox(height: 20),
              Text(userID),
              SizedBox(height: 40),
              Container(
                width: double.infinity,
                padding: EdgeInsets.all(10),
                child: ElevatedButton(
                  onPressed: () {
                    Navigator.pushNamed(context, "/history_login");
                  },
                  child: Row(
                    children: [
                      Padding(
                        padding: const EdgeInsets.only(
                            bottom: 10, top: 10, right: 10),
                        child: Icon(Icons.history), // Biểu tượng
                      ),
                      Text('Lịch sử truy cập', textAlign: TextAlign.left),
                    ],
                  ),
                ),
              ),
              SizedBox(height: 10),
              Container(
                width: double.infinity,
                padding: EdgeInsets.all(10),
                child: ElevatedButton(
                  onPressed: () {
                    showAddAcountPopup(context);
                  },
                  child: Row(
                    children: [
                      Padding(
                        padding: const EdgeInsets.only(
                            bottom: 10, top: 10, right: 10),
                        child: Icon(Icons.lock), // Biểu tượng
                      ),
                      Text('Thay đổi mật khẩu', textAlign: TextAlign.left),
                    ],
                  ),
                ),
              ),
              SizedBox(height: 10),
              Container(
                width: double.infinity,
                padding: EdgeInsets.all(10),
                child: ElevatedButton(
                  onPressed: () {
                    Navigator.pushReplacementNamed(context, '/login');
                  },
                  child: Row(
                    children: [
                      Padding(
                        padding: const EdgeInsets.only(
                            bottom: 10, top: 10, right: 10),
                        child: Icon(Icons.logout), // Biểu tượng
                      ),
                      Text('Đăng xuất', textAlign: TextAlign.left),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
