
func BenchmarkRepbat(b *testing.B){
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
