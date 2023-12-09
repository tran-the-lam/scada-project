import 'package:flutter/material.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:intl/intl.dart';
import 'package:mydemo/utils/http.dart' as http;

class HistoryLoginPage extends StatefulWidget {
  const HistoryLoginPage({super.key, required this.title});
  final String title;

  @override
  State<HistoryLoginPage> createState() => _HistoryList();
}

class _HistoryList extends State<HistoryLoginPage> {
  final List<http.HistoryLogin> items = [];
  
  @override
  void initState() {
    super.initState();
    fetchHistory();
  }

  fetchHistory() async {
    var events = await http.GetHistoryLogin();
    setState(() {
      items.addAll(events);
    });
  }

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
                  // Container(
                  //   padding: const EdgeInsets.all(8.0),
                  //   alignment: Alignment.centerLeft,
                  //   child: Text("Địa chỉ: " + items[index]['ip_address']),
                  // ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Thiết bị: " + items[index].device),
                    alignment: Alignment.centerLeft,
                  ),
                  Container(
                    padding: EdgeInsets.all(8.0),
                    child: Text("Thời gian: " + items[index].createdAt),
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