import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';

enum SensorType { temperature, humidity, pressure }

class ProfilePage extends StatefulWidget {
  const ProfilePage({super.key, required this.title});
  final String title;

  @override
  State<ProfilePage> createState() => _ProfileBody();
}

class _ProfileBody extends State<ProfilePage> {

  String newPwd = '';
  String confirmPwd = '';
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
                    obscureText: true,
                    onChanged: (value) {
                      setState(() {
                        newPwd = value;
                      });
                    },
                    decoration: InputDecoration(hintText: 'Nhập mật khẩu mới'),
                  ),
                  SizedBox(height: 10, width: 10),
                  TextField(
                    obscureText: true,
                    onChanged: (value) {
                      setState(() {
                        confirmPwd = value;
                      });
                    },
                    decoration: InputDecoration(hintText: 'Xác nhận lại mật khẩu'),
                  ),
                  SizedBox(height: 10, width: 10),

                ],
              ),
              actions: [
                ElevatedButton(
                  onPressed: () {
                    
                    if (confirmPwd != newPwd) {
                      Fluttertoast.showToast(
                          msg: "Mật khẩu không khớp",
                          toastLength: Toast.LENGTH_LONG,
                          gravity: ToastGravity.BOTTOM,
                          timeInSecForIosWeb: 1,
                          backgroundColor: Colors.redAccent,
                          textColor: Colors.white,
                          webPosition: "center",
                          fontSize: 16.0);
                    } else {
                      //TODO call API to change password
                      Fluttertoast.showToast(
                          msg: "Đổi mật khẩu thành công",
                          toastLength: Toast.LENGTH_SHORT,
                          gravity: ToastGravity.BOTTOM,
                          timeInSecForIosWeb: 1,
                          backgroundColor: Colors.greenAccent,
                          textColor: Colors.white,
                          webPosition: "center",
                          fontSize: 16.0);
                          print('New password: $newPwd');
                        Navigator.of(context).pop();
                    }
                   
                  },
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
              Text("UserID"),
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
                      padding: const EdgeInsets.only(bottom: 10, top: 10, right: 10),
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
                      padding: const EdgeInsets.only(bottom: 10, top: 10, right: 10),
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
                      padding: const EdgeInsets.only(bottom: 10, top: 10, right: 10),
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
