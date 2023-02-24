<template>
    <div>
        <div>
            <ul>
                <li v-for="(blog, index) of blogs" :key="index" style="list-style-type: none;">
                    <div class="blogs">
                        <div class="blogsLink">
                            <el-link style="font-size:large; font-weight:bolder;">{{
                                blog.title
                            }}</el-link>
                        </div>
                        <div class="blogsCover">
                            <!-- 封面 -->
                        </div>
                        <div class="blogsProfile">
                            <!-- 简介 -->
                        </div>
                    </div>
                </li>
            </ul>
        </div>
        <div style="position: absolute; left: 35%;">
            <el-pagination layout="prev, pager, next" :total="total" :page-size="limit" :current-page.sync="page"
                @current-change="changePage">
            </el-pagination>
        </div>
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
            this.getBlogs()
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

<style>
.blogs {
    height: 150px;
    border:1px solid #dedede;
    border-collapse:collapse;
}
</style>