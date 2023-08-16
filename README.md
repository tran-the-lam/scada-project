Project ứng dụng công nghệ blockchain [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-2.5/index.html) để tăng cường tính bảo mật cho hệ thống `Scada`. 

#### Kiến thức nền tảng:
[SCADA](https://docs.google.com/document/d/10R_ofWSwNjWEZ7i4tid5dPGmwIhwETDxzFmLMU1RotE/edit?usp=sharing) là một hệ thống điều khiển giám sát và thu thập dữ liệu nhằm hỗ trợ con người trong quá trình giám sát và điều khiển từ xa. SCADA là giải pháp tối ưu được các tổ chức công nghiệp lựa chọn nhằm khai thác truy cập dữ liệu và quản lý thiết bị. Người vận hành có thể xem các dữ liệu quan trọng như nhiệt độ, điện áp, mức tiêu hao năng lượng, … và đưa ra các quyết định với sự thay đổi của từng mức dữ liệu tương ứng. Vì vậy bảo mật dữ liệu trong hệ thống SCADA là rất quan trọng.Hiện nay dữ liệu được lưu trên hệ thống SCADA thường được lưu vào các hệ quản trị cơ sở dữ liệu nên sẽ dễ dàng bị tin tặc tấn công hoặc dễ dàng bị thay đổi để tạo ra sự sai lệch, sẽ dẫn đến những quyết định không đúng đắn của người vận hành và gây ra những hậu quả nặng nề.

Chuỗi khối là công nghệ tạo ra hệ thống xác thực mạng ngang hàng, nhằm loại bỏ các bên trung gian thứ ba, tăng cường an ninh, an toàn và minh bạch cũng như giảm thiểu lỗi do con người gây ra. Chuỗi khối được thiết kế để chống lại sự thay đổi của dữ liệu. Dữ liệu đã được mạng lưới chấp nhận thì khó có cách nào có thể thay đổi được. Ứng dụng chuỗi khối vào các hệ thống có các bên tham gia là rất lớn. Trong đề tài này, chúng tôi sẽ ứng dụng chuỗi khối để tăng cường bảo mật dữ liệu cho hệ thống SCADA. Các thay đổi trong hệ thống sẽ được lưu và xem lại. Từ đó, giúp cho việc tìm ra nguyên nhân, khắc phục lỗi khi sự cố xảy ra được dễ dàng.

__Bố cục báo cáo__:
<ol>
  <li>Giới thiệu về SCADA
  <ol>
  <li>SCADA là gì?</li>
  <li>Tầm quan trọng của dữ liệu và thực trạng?</li>
  </ol>
  </li>
  
  <li>Giới thiệu về chuỗi khối
  <ol>
  <li>Chuỗi khối là gì?</li>
  <li>Dữ liệu trên chuỗi khối?</li>
  </ol>
  </li>
  
  <li>Nền tảng <a href="https://docs.google.com/document/d/1OgtoTqUcE656rH7cmA90lOfzGgrw3QyokLfaWRzNBBM/edit?usp=sharing">Hyperledger</a>
    <ol>
      <li>Hyperledger là gì?</li>
      <li>Xây dựng các node trên Hyperledger</li>
      <li>Lựa chọn dữ liệu đưa vào chuỗi khối</li>
      <li>Truy xuất dữ liệu</li>
    </ol>
  </li>
  <li>Thực nghiệm</li>
  <li>Tổng kết</li>
</ol>