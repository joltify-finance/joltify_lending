package keeper

//func TestSeekCorrectPayment(t *testing.T) {
//
//	//senario 1: only one borrow record and the payment 0
//	t1 := time.Now()
//	borrowDetails := []types.BorrowDetail{
//		types.BorrowDetail{BorrowedAmount: sdk.NewCoin("demo", sdk.NewIntFromUint64(10000)), TimeStamp: t1},
//	}
//	testPayment := types.PaymentItem{PaymentTime: t1, PaymentAmount: sdk.NewCoin("demo", sdk.NewIntFromUint64(0))}
//	borrowAmount := seekCorrectBorrow(borrowDetails, &testPayment)
//	require.True(t, borrowAmount.IsEqual(sdk.NewCoin("demo", sdk.NewIntFromUint64(10000))))
//
//	t2 := t1.Add(time.Hour)
//	testPayment.PaymentTime = t2
//
//	borrowAmount = seekCorrectBorrow(borrowDetails, &testPayment)
//	require.True(t, borrowAmount.IsEqual(sdk.NewCoin("demo", sdk.NewIntFromUint64(10000))))
//
//	t3 := t2.Add(time.Hour)
//	borrowDetails = append(borrowDetails, types.BorrowDetail{BorrowedAmount: sdk.NewCoin("demo", sdk.NewIntFromUint64(11000)), TimeStamp: t3})
//
//	borrowDetails = append(borrowDetails, types.BorrowDetail{BorrowedAmount: sdk.NewCoin("demo", sdk.NewIntFromUint64(12000)), TimeStamp: t3.Add(time.Hour)})
//
//	borrowAmount = seekCorrectBorrow(borrowDetails, &testPayment)
//	require.True(t, borrowAmount.IsEqual(sdk.NewCoin("demo", sdk.NewIntFromUint64(10000))))
//
//	baseTime := testPayment.GetPaymentTime()
//	for i := 0; i < 60; i++ {
//		testPayment.PaymentTime = baseTime.Add(time.Minute * time.Duration(i))
//		borrowAmount = seekCorrectBorrow(borrowDetails, &testPayment)
//		require.True(t, borrowAmount.IsEqual(sdk.NewCoin("demo", sdk.NewIntFromUint64(10000))))
//	}
//
//	for i := 61; i < 120; i++ {
//		testPayment.PaymentTime = baseTime.Add(time.Minute * time.Duration(i))
//		borrowAmount = seekCorrectBorrow(borrowDetails, &testPayment)
//		require.True(t, borrowAmount.IsEqual(sdk.NewCoin("demo", sdk.NewIntFromUint64(11000))))
//	}
//
//	for i := 121; i < 200; i++ {
//		testPayment.PaymentTime = baseTime.Add(time.Minute * time.Duration(i))
//		borrowAmount = seekCorrectBorrow(borrowDetails, &testPayment)
//		require.True(t, borrowAmount.IsEqual(sdk.NewCoin("demo", sdk.NewIntFromUint64(12000))))
//	}
//
//}
