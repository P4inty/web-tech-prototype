<script lang="ts">
    import { FileApi } from "../api/file";
    import { onMount } from "svelte";
    import type { CustomFile } from "../types/file.type";

    let files: CustomFile[] = [];
    let searchResults: CustomFile[] = [];
    let query = "";

    onMount(() => {
        FileApi.getFiles().then((res) => {
            files = res.data.files;
        });
    });

    function search() {
        const q = query.toLowerCase();
        searchResults = files.filter((file) => {
            if (file.CreatedAt.toLowerCase().includes(q)) return true;
            if (file.Description.toLowerCase().includes(q)) return true;
            if (file.ID.toString() === q) return true;
            if (file.Tags.some((tag) => tag.Key.toLowerCase().includes(q)))
                return true;
        });
    }
</script>

<div class="m-4">
    <div class="flex justify-end">
        <div class="flex flex-col w-full md:w-1/4 mb-4">
            <label for="search">Search</label>
            <input
                type="text"
                class="input"
                id="search"
                on:keyup={search}
                bind:value={query}
            />
        </div>
    </div>
    <div class="overflow-x-auto">
        <table class="table-auto w-full">
            <thead class="text-left">
                <tr>
                    <th class="table-header">ID</th>
                    <th class="table-header">Filename</th>
                    <th class="table-header">Description</th>
                    <th class="table-header">Creation</th>
                    <th class="table-header">Tags</th>
                    <th class="table-header" />
                </tr>
            </thead>
            <tbody>
                {#each query.length > 0 ? searchResults : files as file}
                    <tr class="table-row">
                        <td class="table-entry">{file.ID}</td>
                        <td class="table-entry">{file.Name}</td>
                        <td class="table-entry">{file.Description}</td>
                        <td class="table-entry"
                            >{new Date(file.CreatedAt).toLocaleDateString()}</td
                        >
                        <td class="space-y-1 items-stretch">
                            {#each file.Tags as tag}
                                <button
                                    class="bg-primary-dark px-2 hover:text-secondary hover:bg-primary mr-1"
                                    on:click={() => {
                                        query = tag.Key;
                                        search();
                                    }}>{tag.Key}</button
                                >
                            {/each}
                        </td>
                        <td>
                            <a
                                class="btn border-y-0 block text-center"
                                href={import.meta.env.VITE_API_URL +
                                    `/download/${file.Uri}`}
                                download>Download</a
                            >
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </div>
</div>
