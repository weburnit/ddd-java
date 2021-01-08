# Xây dựng hệ thống quản trị Reward(Offer) cơ bản

Hệ thống có 2 nhóm user cơ bản: VDS Staff User và Partner User

## Operation
* Reward đc cung cấp bởi Partner và các Partner chỉ định đc điều kiện (Policy) là Demographic.
	VD: CGV muốn gửi Voucher cho các User nữ nhân ngày 8/3

* Reward đc duyệt bởi VDS Staff, sau khi duyệt thành công, Voucher này sẽ tồn tại trong kho Voucher để cấp phát cho các User của ViettelPayl


## Policy
* Reward được chỉ định dưới 2 dạng khác nhau:
* Voucher: có code để Redeem
* Giảm giá: giảm theo %X và tối đa không quá Y VNĐ

* Reward đc chỉ định bởi các Policy độc lập dựa vào DemographicPolicy(Nhân khẩu học) và EventPolicy
	1. Demographic: với các Policy tương ứng về nhân khẩu học theo
		a. Age
		b. Gender****
		VD: Reward là Voucher vé xem phim Batman với nhóm khách hàng Nữ, dưới 22 tuổi.
			Khi đó, hệ thống chỉ cấp phát cho những user thoã mãn các Policy trên.

	2. EventPolicy
		a. event_name: tên của các event cụ thể
		b. timeStamp: thời gian event thoã mãn
		VD: Cấp phát Voucher giảm 20% cho các khách hàng event_name = thanh_toan_khoan_vay vào ngày thứ 7 hàng tuần.
	3. TimeBasedPolicy:
        ** Ngày cố định(absolute)
        ** Ngày trong tuần

* Có thể tồn tại Các Reward mà tổ hợp cả 2 điều kiện ở trên (AND)


## Event Driven Flow
* Với các sự kiện TriggerEvent của Users từ Kafka, hệ thống sẽ process các event tương ứng để tính toán các Reward phù hợp.
* Khi các User nhận đc Reward, hệ thống gửi email cho khách hàng thông tin voucher.

```shell script
# Trigger Event class
TriggerEvent: đc stream vào Kafka và xử lý để tính toán Reward tương ứng
	event_name: string
	properties: HashMap<String, Object>
	user: Profile
```
# Yêu cầu đề bài
* Hệ thống cho phép 2 nhóm Users: Partner và VDS Staff login vào hệ thống
* Partner có thể khởi tạo Reward với Demographic Policy(age, gender) và TimeBasedPolicy
* VDS Staff có thể thiết lập thêm EventPolicy cho các Reward tương ứng.
* VDS Staff approve hoặc Reject (statemachine)
* Khi một sự kiện TriggerEvent xảy ra (thông qua Kafka), hệ thống xử lý (Consumer) luồng event (stream) này để tính toán chấp nhận hay không việc Tạo ra OfferItem cho Profile tương ứng của Event.
* Với các User xác định Fraud, cần lưu lại(KafkaAllocationEmitter) trên Kafka Event
* Với các User xác định số lần nhận reward trong ngày sẽ lưu vào Database(SQLAllocationEmitter)
* Mỗi lần OfferItem đc tạo ra, cần send Email đến cho Partner.email
