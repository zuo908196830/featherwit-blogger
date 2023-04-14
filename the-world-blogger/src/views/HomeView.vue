<template>
    <div>
        <div>
            <ul>
                <li v-for="(blog, index) of blogs" :key="index" style="list-style-type: none;">
                    <div class="blogs">
                        <div class="blogsLink">
                            <el-link style="font-size:large; font-weight:bolder;"
                                @click="getBlogById(blog)">{{
                                    blog.title
                                }}</el-link>
                        </div>
                        <div class="blogsCover" v-if="blog.cover">
                            <!-- 封面 -->
                            <el-image style="width: 150px;height: 125px" :src="blog.cover" :fit="fill "></el-image>
                        </div>
                        <div class="blogsProfile" v-if="blog.cover">
                            <!-- 简介 -->
                            <div>{{ blog.profile }}</div>
                        </div>
                        <div v-if="!blog.cover">
                            <!-- 简介 -->
                            <div>{{ blog.profile }}</div>
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
            loading: false,
            searchMsg: {}
        }
    },
    methods: {
        changePage() {
            localStorage.setItem("page", this.page)
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
            axios.post('/api/blog/' + this.limit + '/' + (this.limit) * (this.page - 1), this.searchMsg).then(res => {
                if (res.data.code === 0) {
                    this.blogs = res.data.data
                }
            }).catch(() => {
                this.$message({
                    message: '服务器错误',
                    type: 'error'
                })
            })
        },
        getBlogById(blog) {
            this.$router.push("/blog?id=" + blog.id)
        }
    },
    created() {
        var p = localStorage.getItem('page')
        if (p) {
            this.page = p
        }
        this.getTotal()
        this.getBlogs()
    },
}
</script>

<style>
.blogs {
    height: 150px;
    border: 1px solid #dedede;
    border-collapse: collapse;
}

.blogsProfile {
    position: absolute;
    left: 210px;
}

.blogsCover {
    position: absolute;
    left: 50px;
}
</style>