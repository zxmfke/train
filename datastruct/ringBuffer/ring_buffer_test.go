package ringBuffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	rb := NewRingBuffer(10)
	t.Logf("%s", rb)

	_ = rb.Set(10, 111)
	t.Logf("%s", rb)
}

func TestRingBuffer_Set(t *testing.T) {
	root := NewRingBuffer(3)

	cases := []struct {
		name      string
		setKey    int
		setValue  interface{}
		wantErr   error
		wantTotal int
		wantHead  int
		wantGet   bool
	}{
		{
			name:      "插入第一个 key",
			setKey:    1,
			setValue:  1,
			wantErr:   nil,
			wantTotal: 1,
			wantHead:  1,
		}, {
			name:      "插入第二个 key",
			setKey:    2,
			setValue:  2,
			wantErr:   nil,
			wantTotal: 2,
			wantHead:  2,
		}, {
			name:      "插入相同的 key",
			setKey:    3,
			setValue:  3,
			wantErr:   nil,
			wantTotal: 3,
			wantHead:  0,
		}, {
			name:     "插入一个满的 ring buffer",
			setKey:   4,
			setValue: 4,
			wantErr:  errFull,
		}, {
			name:      "取出一个",
			wantGet:   true,
			wantTotal: 2,
		}, {
			name:      "满了之后再插入 key",
			setKey:    5,
			setValue:  5,
			wantErr:   nil,
			wantTotal: 3,
			wantHead:  1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			if tt.wantGet {
				_, _ = root.Get()
				t.Logf("%s", root)
				return
			}

			err := root.Set(tt.setKey, tt.setValue)
			if err != nil {
				if tt.wantErr != nil {
					assert.Equal(t, tt.wantErr, err)
					return
				}
			}

			t.Logf("%s", root)
			assert.Equal(t, tt.wantTotal, root.total)
			assert.Equal(t, tt.wantHead, root.head)
		})
	}
}

func TestRingBuffer_Get(t *testing.T) {
	root := NewRingBuffer(3)

	_ = root.Set(1, 1)
	_ = root.Set(2, 2)
	_ = root.Set(3, 3)

	t.Logf("%s", root)

	cases := []struct {
		name      string
		wantValue int
		wantErr   error
		wantTotal int
		wantTail  int
		wantSet   bool
	}{
		{
			name:      "取出第一个",
			wantValue: 1,
			wantErr:   nil,
			wantTotal: 2,
			wantTail:  1,
		}, {
			name:      "取出第二个",
			wantValue: 2,
			wantErr:   nil,
			wantTotal: 1,
			wantTail:  2,
		}, {
			name:      "取出第三个",
			wantValue: 3,
			wantErr:   nil,
			wantTotal: 0,
			wantTail:  0,
		}, {
			name:      "从一个空的ring buffer取",
			wantValue: 4,
			wantErr:   errEmpty,
		}, {
			name:      "能继续写入",
			wantValue: 4,
			wantErr:   nil,
			wantTotal: 1,
			wantSet:   true,
			wantTail:  1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantSet {
				err := root.Set(1, 1)
				if err != nil {
					assert.Equal(t, tt.wantErr, err)
					return
				}

				assert.Equal(t, tt.wantTotal, root.total)
				return
			}

			v, err := root.Get()
			if err != nil {
				if tt.wantErr != nil {
					assert.Equal(t, tt.wantErr, err)
					return
				}
			}

			t.Logf("%s", root)
			assert.Equal(t, tt.wantTotal, root.total)
			assert.Equal(t, tt.wantTail, root.tail)
			assert.Equal(t, tt.wantValue, v.(int))
		})
	}
}
