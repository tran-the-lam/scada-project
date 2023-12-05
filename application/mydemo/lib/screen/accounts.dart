import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';
import 'package:mydemo/utils/http.dart' as http;

enum SensorType { temperature, humidity, pressure }

class AccountPage extends StatefulWidget {
  const AccountPage({super.key, required this.title});
  final String title;

  @override
  State<AccountPage> createState() => _AccountList();
}

class _AccountList extends State<AccountPage> {
  final List<http.User> items = [];

  DateFormat formatter = DateFormat('yyyy/MM/dd');
  DateTime selectedStartDate = DateTime.now();
  DateTime selectedEndDate = DateTime.now();
  List<SensorType> selectedSensorTypes = [];
  TextEditingController _userIDController = TextEditingController();
  List<String> options = ['manager', 'employee'];
  String _role = 'manager';

  @override
  void initState() {
    super.initState();
    fetchAllUser();
  }

  void fetchAllUser() async {
    final users = await http.getAllUser();
    setState(() {
      items.clear();
      items.addAll(users);
    });
  }

  Future<void> addAccountOnPress() async {
    _role = _role.trim();

    if (_userIDController.text.isEmpty) {
      Fluttertoast.showToast(
        msg: "Vui lòng nhập đầy đủ thông tin",
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.CENTER,
        timeInSecForIosWeb: 1,
      );
      return;
    }

    final result = await http.addUser(_userIDController.text, _role);
    print(result);
    if (result == "success") {
      setState(() {
        items.insert(0, http.User(userID: _userIDController.text, role: _role, status: 1));
      });
      
      Fluttertoast.showToast(
        msg: "Thêm tài khoản thành công",
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.CENTER,
        timeInSecForIosWeb: 1,
      );
      
    } else {
      Fluttertoast.showToast(
        msg: "Thêm tài khoản thất bại",
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.CENTER,
        timeInSecForIosWeb: 1,
      );
    }
    Navigator.of(context).pop();
  }

  void showAddAccountPopup(BuildContext context) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return StatefulBuilder(
          builder: (context, setState) {
            return AlertDialog(
              title: Text('Tạo tài khoản',
                  style: TextStyle(fontWeight: FontWeight.bold),
                  textAlign: TextAlign.center),
              content: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  TextField(
                    controller: _userIDController,
                    decoration: InputDecoration(hintText: 'Nhập mã nhân viên'),
                  ),
                  // SizedBox(height: 10, width: 10),
                  // TextField(
                  //   onChanged: (value) {
                  //     setState(() {
                  //       pwd = value;
                  //     });
                  //   },
                  //   decoration: InputDecoration(hintText: 'Nhập mật khẩu'),
                  // ),
                  SizedBox(height: 10, width: 10),
                  Row(children: [
                    Text('Vai trò:',
                        style: TextStyle(fontWeight: FontWeight.bold)),
                    SizedBox(height: 10, width: 10),
                    DropdownButton(
                      value: _role,
                      items: options.map((String option) {
                        return DropdownMenuItem(
                          value: option,
                          child: Text(option),
                        );
                      }).toList(),
                      onChanged: (String? newValue) {
                        setState(() {
                          _role = newValue!;
                        });
                      },
                    ),
                  ]),
                ],
              ),
              actions: [
                ElevatedButton(
                  onPressed: addAccountOnPress,
                  child: Text('Thêm'),
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
    const title = 'Users';
    return MaterialApp(
      title: title,
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: const Text(title),
          actions: [
            IconButton(
              icon: Icon(Icons.add_circle_outline_sharp),
              onPressed: () {
                showAddAccountPopup(context);
              },
            ),
          ],
        ),
        body: ListView.builder(
          itemCount: items.length,
          itemBuilder: (context, index) {
            return Card(
              child: Column(
                children: [
                  Container(
                    child: Text("Mã nhân viên: " + items[index].userID),
                    padding: EdgeInsets.all(8.0),
                    alignment: Alignment.centerLeft,
                  ),
                  // Container(
                  //   child: Text("Vai trò: " + items[index].role),
                  //   padding: EdgeInsets.all(8.0),
                  //   alignment: Alignment.centerLeft,
                  // ),
                  Container(
                    // child: Text("Thời gian tạo: " + items[index]['created_at']),
                    padding: EdgeInsets.only(left: 8.0, right: 8.0),
                    alignment: Alignment.centerLeft,
                    child: Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Text("Vai trò: " + items[index].role),
                        SizedBox(
                            width:
                                350), // Khoảng cách giữa văn bản và PopupMenuButton
                        PopupMenuButton(
                          itemBuilder: (context) {
                            return [
                              PopupMenuItem(
                                value: 'reset_pwd',
                                child: Text('Reset mật khẩu'),
                              ),
                              // PopupMenuItem(
                              //   value: 'delete',
                              //   child: Text('Xóa tài khoản'),
                              // )
                            ];
                          },
                          onSelected: (String value) {
                            print('You Click on po up menu item ${value}');
                          },
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            );
          },
        ),
      ),
    );
  }
}
