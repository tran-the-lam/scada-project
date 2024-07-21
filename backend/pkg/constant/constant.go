package constant

const (
	LOCAL_USER_ID    = "user_id"    // Lưu mã người dùng trong phiên
	LOCAL_USER_ROLE  = "user_role"  // Lưu vai trò người dùng trong phiên
	LOCAL_SENSOR_ID  = "sensor_id"  // Lưu mã cảm biến trong phiên
	LOCAL_DEVICE_ID  = "device_id"  // Lưu mã thiết bị trong phiên
	LOCAL_IP_ADDR    = "ip_addr"    // Lưu địa chỉ IP trong phiên
	LOCAL_USER_AGENT = "user_agent" // Lưu thông tin User-Agent trong phiên
	TOKEN_SECRET     = "secret"     // Mã bí mật dùng để tạo mã JWT

	API_KEY = "scada-api-key" // Khóa API

	TEMPERATURE_SENSOR = "temperature" // Loại cảm biến nhiệt độ
	HUMIDITY_SENSOR    = "humidity"    // Loại cảm biến độ ẩm

	ADMIN_ROLE    = "admin"    // Vai trò quản trị viên
	MANAGER_ROLE  = "manager"  // Vai trò quản lý
	EMPLOYEE_ROLE = "employee" // Vai trò nhân viên

	// Các hàm trong chuỗi mã
	SMC_FUNC_VERIFY_USER             = "VerifyUser"            // Hàm kiểm tra người dùng
	SMC_FUNC_SAVE_LOGIN              = "SaveLoginInfo"         // Hàm lưu thông tin đăng nhập
	SMC_FUNC_QUERY_KEY               = "QueryKey"              // Hàm truy vấn khóa
	SMC_FUNC_ADD_USER                = "AddUser"               // Hàm thêm người dùng
	SMC_FUNC_UPDATE_PWD              = "UpdatePassword"        // Hàm cập nhật mật khẩu
	SMC_FUNC_GET_TRANSACTION_HISTORY = "GetTransactionHistory" // Hàm lấy lịch sử giao dịch
	SMC_FUNC_ADD_EVENT               = "AddEvent"              // Hàm thêm sự kiện
	SMC_FUNC_GET_ALL_EVENTS          = "GetAllEvents"          // Hàm lấy tất cả sự kiện
	SMC_FUNC_GET_EVENTS_BY_KEY       = "GetEventsByKey"        // Hàm lấy sự kiện theo khóa
	SMC_FUNC_RESET_PWD               = "ResetPassword"         // Hàm đặt lại mật khẩu
	SMC_FUNC_DELETE_USER             = "DeleteUser"            // Hàm xóa người dùng
	SMC_FUNC_GET_ALL_USERS           = "GetAllUsers"           // Hàm lấy tất cả người dùng
)
