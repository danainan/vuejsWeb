
<template>
    <div class="relative flex flex-col justify-center h-screen overflow-hidden">
        <div class="w-full p-6 m-auto bg-white rounded-md shadow-md lg:max-w-lg">
            <h1 class="text-3xl font-semibold text-center text-purple-700">Panda</h1>
            <form class="space-y-4">
                <div>
                    <label class="label">
                        <span class="text-base label-text">Username</span>
                    </label>
                    <input type="text" placeholder="Username" class="w-full input input-bordered input-primary" />
                </div>
                <div>
                    <label class="label">
                        <span class="text-base label-text">Email</span>
                    </label>
                    <input type="email" placeholder="Email Address" class="w-full input input-bordered input-primary" />
                </div>
                <div>
                    <label class="label">
                        <span class="text-base label-text">Password</span>
                    </label>
                    <input type="password" placeholder="Enter Password"
                        class="w-full input input-bordered input-primary" />
                </div>
                <div>
                    <button class="btn btn-primary">Sign up</button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup>
import 'axios'
import {ref} from 'vue'

const user = ref({
    username : '',
    email : '',
    password : ''

})

const register = () =>

</script>
