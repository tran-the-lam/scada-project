import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';

class HistoryLoginPage extends StatefulWidget {
  const HistoryLoginPage({super.key, required this.title});
  final String title;

  @override
  State<HistoryLoginPage> createState() => _HistoryList();
}

class _HistoryList extends State<HistoryLoginPage> {
  final List<Map<String, dynamic>> items = [
    {
      'ip_address': '192.168.0.1',
      'device': 'IPhone 12 Pro Max',
      'created_at': '2023-10-10 10:10:10'
    },
    {
      'ip_address': '192.168.1.1',
      'device': 'Huawei P40 Pro',
      'created_at': "2023-10-08 09:10:10"
    },
    {
      'ip_address': '23.123.432.12',
      'device': 'Samsung Galaxy S21',
      'created_at': "2023-10-08 08:10:10"
    },
    {
      'ip_address': '123.12.32.12',
      'device': 'Macbook Air 2020',
      'created_at': "2023-10-08 07:10:10"
    },
    {
      'ip_address': '231.12.342.12',
      'device': 'Macbook Pro 2020',
      'created_at': "2023-10-08 06:10:10"
    },
    {
      'ip_address': '231.341.12.32',
      'device': 'BPhone 4',
      'created_at': "2023-10-08 05:10:10"
    },
    
  ];

  @override
  Widget build(BuildContext context) {
    const title = 'Lịch sử đăng nhập';

    return MaterialApp(
      title: title,
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: false,
          title: const Text(title),
          leading: IconButton(
          icon: Icon(Icons.arrow_back),
          onPressed: () {
            Navigator.pop(context);
          },
        ),
        ),
        body: ListView.builder(
          itemCount: items.length,
          itemBuilder: (context, index) {
            return Card(
              child: Column(
                children: [
                  Container(
                    padding: const EdgeInsets.all(8.0),
                    alignment: Alignment.centerLeft,
                    child: Text("Địa chỉ: " + items[index]['ip_address']),
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Thiết bị: " + items[index]['device']),
                    alignment: Alignment.centerLeft,
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Thời gian: " + items[index]['created_at']),
                    alignment: Alignment.centerLeft,
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