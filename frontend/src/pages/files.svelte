<script lang="ts">
    import { FileApi } from "../api/file";
    import { onMount } from "svelte";
    import type { CustomFile } from "../types/file.type";

    let files: CustomFile[] = [];

    onMount(() => {
        FileApi.getFiles().then((res) => {
            files = res.data.files;
        });
    });
</script>

<div class="m-4">
    <table class="table-auto w-full">
        <thead class="text-left">
            <tr>
                <th class="table-header">ID</th>
                <th class="table-header">Filename</th>
                <th class="table-header">Description</th>
                <th class="table-header">Tags</th>
                <th class="table-header" />
            </tr>
        </thead>
        <tbody>
            {#each files as file}
                <tr class="table-row">
                    <td class="table-entry">{file.ID}</td>
                    <td class="table-entry">{file.Name}</td>
                    <td class="table-entry">{file.Description}</td>
                    <td class="space-x-1 items-stretch">
                        {#each file.Tags as tag}
                            <span class="bg-primary-dark px-2">{tag.Key}</span>
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
