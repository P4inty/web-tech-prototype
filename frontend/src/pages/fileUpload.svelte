<script type="ts">
    import type { MetaData } from "../types/file.type";
    import Dropzone from "../lib/dropzone.svelte";
    import { FileApi } from "../api/file";
    import { ViewManager } from "../stores/view";
    import { Views } from "../enums/views";

    let file: File = null;
    let meta: MetaData = {
        Name: "",
        Description: "",
        Tags: [],
    };
    let currentTag: String = "";

    function addTag(e: any) {
        if ([32, 13].includes(e.keyCode)) {
            meta.Tags = [...meta.Tags, { Key: currentTag }];
            currentTag = "";
        }
    }

    function removeTag(i: number) {
        meta.Tags.splice(i, 1);
        meta.Tags = [...meta.Tags];
    }

    function upload() {
        FileApi.uploadMetadata(meta).then((res) =>
            FileApi.uploadFile(file, res.data.uri).then(() =>
                ViewManager.goTo(Views.FILES)
            )
        );
    }
</script>

<div class="m-4">
    {#if !file}
        <Dropzone bind:file />
    {:else}
        <p class="text-2xl mt-4">Additional File Information</p>
        <form class="flex flex-col space-y-4" on:submit|preventDefault={upload}>
            <div class="flex flex-col mt-4">
                <label for="name">Name</label>
                <input
                    type="text"
                    name="name"
                    class="input"
                    required
                    bind:value={meta.Name}
                />
            </div>
            <div class="flex flex-col mt-4">
                <label for="contentDescription">Content Description</label>
                <textarea
                    name="contentDescription"
                    cols="30"
                    rows="10"
                    class="input"
                    required
                    bind:value={meta.Description}
                />
            </div>
            <div class="flex flex-col mt-4">
                <label for="tags">Tags</label>
                <div class="input flex">
                    <div class="flex space-x-1">
                        {#each meta.Tags as tag, i}
                            <span class="bg-primary-dark px-2 flex items-center"
                                >{tag.Key}
                                <button
                                    class="text-primary hover:text-secondary pl-2"
                                    on:click={() => removeTag(i)}>x</button
                                ></span
                            >
                        {/each}
                    </div>
                    <input
                        type="text"
                        name="tags"
                        class="bg-transparent w-full outline-none pl-2"
                        on:keydown={(e) => addTag(e)}
                        bind:value={currentTag}
                        placeholder="Add with (Space)"
                    />
                </div>
            </div>
            <div class="flex space-x-4">
                <button
                    class="btn w-1/4 text-primary"
                    type="button"
                    on:click={() => (file = null)}>Cancel</button
                >
                <button class="btn" type="submit">Upload</button>
            </div>
        </form>
    {/if}
</div>
