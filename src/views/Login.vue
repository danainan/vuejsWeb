<script setup>
import { RouterLink, RouterView} from 'vue-router'

</script>
<template>
    <div class="relative flex flex-col justify-center h-screen overflow-hidden">
        <div class="w-full p-6 m-auto bg-white rounded-md shadow-md lg:max-w-lg">
            <h1 class="text-3xl font-semibold text-center text-purple-700">Panda</h1>
            <form class="space-y-4">
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
                <a href="#" class="text-xs text-gray-600 hover:underline hover:text-blue-600">Forget Password?</a>
                <div>
                    <button class="btn btn-primary">Login</button>
                    <span>
                        <RouterLink to="register"><a class="pl-5 underline underline-offset-1 hover:text-purple-400/50" href="">sign up</a></RouterLink>
                    </span>
                </div>
            </form>
        </div>
    </div>
</template>