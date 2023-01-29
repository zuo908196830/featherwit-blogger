<template>
    <div>
        <div>
            <el-link v-for="(blog, index) of blogs" :key="index">{{ blog.title }}</el-link>
        </div>
        <el-pagination
            layout="prev, pager, next"
            :total="total"
            :page-size="limit"
            :current-page.sync="page"
            @current-change="changePage">
        </el-pagination>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            blogs: [],
            page: 1,
            limit: 10,
            total: 90,
            loading: false
        }
    },
    methods: {
        changePage() {
            alert(this.page)
        },
        getTotal() {
            axios.get('/api/blog/count').then(res => {
                if (res.data.code === 0) {
                    this.total = res.data.data.total
                }
            }).catch(() => {
                this.$message({
                    message: '服务器错误',
                    type: 'error'
                })
            })
        },
        getBlogs() {
            axios.get('/api/blog/' + this.limit + '/' + (this.limit) * (this.page - 1)).then(res => {
                if (res.data.code === 0) {
                    this.blogs = res.data.data
                }
            }).catch(() => {

            })
        }
    },
    created() {
        this.getTotal()
        this.getBlogs()
    },
}
</script>