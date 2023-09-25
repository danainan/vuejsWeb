<script setup>
import { RouterLink, RouterView} from 'vue-router'

</script>
<template>
    <div class="navbar bg-white-400">
      <div class="flex-1">
        <RouterLink to="home" class="btn btn-ghost normal-case text-xl">Panda</RouterLink>
      </div>
      <div class="flex-none gap-2">
        <div class="form-control">
          <input type="text" placeholder="Search" class="input input-bordered w-24 md:w-auto" />
        </div>
        <div class="dropdown dropdown-end">
          <label tabindex="0" class="btn btn-ghost btn-circle avatar">
            <div class="w-10 rounded-full">
              <img src="https://avatars.githubusercontent.com/u/71175110?v=4" />
            </div>
          </label>
          <ul tabindex="0" class="mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52">
            <li><RouterLink to="login">login</RouterLink></li>
          </ul>
        </div>
      </div>
    </div>
    <RouterView/>
</template>