package utils

import (
	"fmt"
	"github.com/unknwon/com"
	"strings"
)

// 格式化 页码和页容量
func FormatPager(page, size string) (int, int) {
	newPage := com.StrTo(page).MustInt()
	if newPage < 1 {
		newPage = 1
	}
	newSize := com.StrTo(size).MustInt()
	if newSize < 1 {
		newSize = 10
	}
	if newSize > 20 {
		newSize = 20
	}
	return newPage, newSize
}

func PageHtml(url string, page, size, count int) string {
	pageCount := 1
	if count%size == 0 {
		pageCount = count / size
	} else {
		pageCount = count/size + 1
	}
	res := `<div class="page_box">`

	if pageCount <= 1 {
		res += `<span class="first">首页</span>`
	} else {
		res += `<a class="first" href="` + url + `">首页</a>`
	}

	if page > 1 {
		res += `<a class="prev" href="` + fmt.Sprintf("%s?page=%d&size=%d", url, page-1, size) + `">上一页</a>`
	} else {
		res += `<span class="prev">上一页</span>`
	}

	res += `<span class="show">` + fmt.Sprintf("%d/%d", page, pageCount) + `</span>`

	if pageCount > page {
		res += `<a class="next" href="` + fmt.Sprintf("%s?page=%d&size=%d", url, page+1, size) + `">下一页</a>`
	} else {
		res += `<span class="next">下一页</span>`
	}

	if page < pageCount {
		res += `<a class="last" href="` + fmt.Sprintf("%s?page=%d&size=%d", url, pageCount, size) + `">尾页</a>`
	} else {
		res += `<span class="last">尾页</span>`
	}

	res += `</div>`
	return res
}

func PageWebHtml(url string, page, size, count int) string {
	pageCount := 1
	if count%size == 0 {
		pageCount = count / size
	} else {
		pageCount = count/size + 1
	}
	res := `<div class="page_box">`

	if pageCount <= 1 {
		res += `<span class="first">首页</span>`
	} else {
		res += `<a class="first" href="` + strings.Replace(url, "{page}", "1", 1) + `">首页</a>`
	}

	if page > 1 {
		res += `<a class="prev" href="` + strings.Replace(url, "{page}", string(page-1), 1) + `">上一页</a>`
	} else {
		res += `<span class="prev">上一页</span>`
	}

	res += `<span class="show">` + fmt.Sprintf("%d/%d", page, pageCount) + `</span>`

	if pageCount > page {
		res += `<a class="next" href="` + strings.Replace(url, "{page}", string(page+1), 1) + `">下一页</a>`
	} else {
		res += `<span class="next">下一页</span>`
	}

	if page < pageCount {
		res += `<a class="last" href="` + strings.Replace(url, "{page}", string(pageCount), 1) + `">尾页</a>`
	} else {
		res += `<span class="last">尾页</span>`
	}

	res += `</div>`
	return res
}
