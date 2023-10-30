package unitTest

// https://note.com/rescuenow_hr/n/n9ed7caf4646d
import (
	"testing"
)

func TestUser_GetMessage(t *testing.T) {

	type fields struct {
	    Name string
    	Age  int
    }

    type args struct {
        message string
    }

	tests := []struct {
        name   string
        fields fields
        args   args
		want   string
	}{
        {"テスト1", fields{"山田", 22}, args{"こんにちは"}, "山田(22)さん, こんにちは"},
        {"テスト2", fields{"佐藤", 26}, args{"はじめまして"}, "佐藤(26)さん, はじめまして"},
    }

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
			t.Logf("t.Logf  %s \n", tt.name)
            u := &User{
                Name: tt.fields.Name,
                Age:  tt.fields.Age,
            }
            if got := u.GetMessage(tt.args.message); got != tt.want {
                t.Errorf("GetMessage() = %v, want %v", got, tt.want)
            }
        })
    }
}