import { useAxiosPrivate } from "@/app/hooks/useAxios";
import { Channel, useFetchChannels } from "@/app/hooks/useChannels";
import { WatchedChannel, WatchedChannelTitleRegex, useCreateWatchedChannel, useEditWatchedChannel } from "@/app/hooks/useWatchedChannels";
import { ActionIcon, Button, NumberInput, TextInput, Tooltip, Text, Divider, Checkbox, Select, Title, Box, Group, Grid, MultiSelect } from "@mantine/core";
import { useForm } from "@mantine/form";
import { showNotification } from "@mantine/notifications";
import { IconPlus, IconTrash } from "@tabler/icons-react";
import { useEffect, useState } from "react";
import classes from "./Watched.module.css"
import { getTwitchCategories } from "@/app/hooks/useCategory";

type Props = {
  watchedChannel: WatchedChannel | null
  mode: WatchedChannelEditMode
  handleClose: () => void;
}

export enum WatchedChannelEditMode {
  Create = "create",
  Edit = "edit",
}

const qualityOptions = [
  { label: "Best", value: "best" },
  { label: "720p60", value: "720p60" },
  { label: "480p", value: "480p30" },
  { label: "360p", value: "360p30" },
  { label: "160p", value: "160p30" },
  { label: "audio", value: "audio" }
];

interface SelectOption {
  label: string;
  value: string;
}

const AdminWatchedChannelDrawerContent = ({ watchedChannel, mode, handleClose }: Props) => {
  const axiosPrivate = useAxiosPrivate()
  const [liveTitleRegexes, setLiveTitleRegexes] = useState<WatchedChannelTitleRegex[]>(
    watchedChannel?.edges.title_regex || []
  );

  const [channelSelect, setChannelSelect] = useState<SelectOption[]>([]);

  // Initialize edit watched channel mutation
  const editWatchedChannelMutation = useEditWatchedChannel();

  const form = useForm({
    mode: "controlled",
    initialValues: {
      id: watchedChannel?.id || "",
      watch_live: watchedChannel?.watch_live ?? false,
      watch_vod: watchedChannel?.watch_vod ?? false,
      download_archives: watchedChannel?.download_archives ?? true,
      download_highlights: watchedChannel?.download_highlights ?? true,
      download_uploads: watchedChannel?.download_uploads ?? true,
      resolution: watchedChannel?.resolution || "best",
      archive_chat: watchedChannel?.archive_chat ?? true,
      channel_id: watchedChannel?.edges.channel.id || "",
      render_chat: watchedChannel?.render_chat ?? true,
      download_sub_only: watchedChannel?.download_sub_only ?? false,
      video_age: watchedChannel?.video_age || 0,
      apply_categories_to_live: watchedChannel?.apply_categories_to_live ?? false,
      watch_clips: watchedChannel?.watch_clips ?? false,
      clips_limit: watchedChannel?.clips_limit || 5,
      clips_interval_days: watchedChannel?.clips_interval_days || 7,
      live_title_regexes: [],
      categories: [] as string[],
    },
  })

  useEffect(() => {
    if (!watchedChannel || !watchedChannel.edges.categories) return

    const categories = watchedChannel.edges.categories.map((category) => category.name);
    form.setFieldValue('categories', categories);

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [watchedChannel])

  const [twitchCategoriesLoading, setTwitchCategoriesLoading] = useState(false);
  const [formattedTwitchCategories, setFormattedTwitchCategories] = useState<SelectOption[]>([]);

  const handleGetTwitchCategories = async () => {
    try {
      setTwitchCategoriesLoading(true)

      const categories = await getTwitchCategories()
      if (!categories) return

      const tmpArr = categories.map((category) => ({
        label: category.name,
        value: category.name,
      })).filter((item, index, self) =>
        index === self.findIndex((t) => t.label === item.label)
      );

      setFormattedTwitchCategories(tmpArr);
    } catch (error) {
      console.error(error)
    } finally {
      setTwitchCategoriesLoading(false)
    }
  }

  const createWatchedChannelMutation = useCreateWatchedChannel();

  const handleSubmitForm = async () => {
    const formValues = form.getValues()

    try {
      if (mode === WatchedChannelEditMode.Create) {
        const newWatchedChannel: WatchedChannel = {
          id: "", // Will be generated by backend
          watch_live: formValues.watch_live,
          watch_vod: formValues.watch_vod,
          download_archives: formValues.download_archives,
          download_highlights: formValues.download_highlights,
          download_uploads: formValues.download_uploads,
          resolution: formValues.resolution,
          archive_chat: formValues.archive_chat,
          render_chat: formValues.render_chat,
          download_sub_only: formValues.download_sub_only,
          video_age: formValues.video_age,
          apply_categories_to_live: formValues.apply_categories_to_live,
          watch_clips: formValues.watch_clips,
          clips_limit: formValues.clips_limit,
          clips_interval_days: formValues.clips_interval_days,
          is_live: false, // Default value
          edges: {
            channel: { id: formValues.channel_id } as Channel,
            categories: [],
            title_regex: liveTitleRegexes
          },
          last_live: "",
          updated_at: "",
          created_at: ""
        };

        await createWatchedChannelMutation.mutateAsync({
          axiosPrivate,
          channelId: formValues.channel_id,
          watchedChannel: newWatchedChannel,
          categories: formValues.categories
        });

        showNotification({
          message: "Watched Channel Created",
          color: "green"
        });

        handleClose();
      } else if (mode === WatchedChannelEditMode.Edit && watchedChannel) {
        const updatedWatchedChannel: WatchedChannel = {
          ...watchedChannel,
          watch_live: formValues.watch_live,
          watch_vod: formValues.watch_vod,
          download_archives: formValues.download_archives,
          download_highlights: formValues.download_highlights,
          download_uploads: formValues.download_uploads,
          resolution: formValues.resolution,
          archive_chat: formValues.archive_chat,
          render_chat: formValues.render_chat,
          download_sub_only: formValues.download_sub_only,
          video_age: formValues.video_age,
          apply_categories_to_live: formValues.apply_categories_to_live,
          watch_clips: formValues.watch_clips,
          clips_limit: formValues.clips_limit,
          clips_interval_days: formValues.clips_interval_days,
          edges: {
            ...watchedChannel.edges,
            title_regex: liveTitleRegexes
          }
        };

        await editWatchedChannelMutation.mutateAsync({
          axiosPrivate,
          watchedChannel: updatedWatchedChannel,
          categories: formValues.categories
        });

        showNotification({
          message: "Watched Channel Updated",
          color: "green"
        });

        handleClose();
      }
    } catch (error) {
      console.error(error);
      showNotification({
        message: mode === WatchedChannelEditMode.Create
          ? "Failed to create Watched Channel"
          : "Failed to update Watched Channel",
        color: "red"
      });
    }
  }


  const { data: channels } = useFetchChannels();

  useEffect(() => {
    if (!channels) return;

    const transformedChannels: SelectOption[] = channels.map((channel: Channel) => ({
      label: channel.name,
      value: channel.id,
    }));

    setChannelSelect(transformedChannels);
  }, [channels]);

  return (
    <div>
      <form onSubmit={form.onSubmit(() => {
        handleSubmitForm()
      })}>
        <TextInput
          disabled={true}
          label="ID"
          placeholder="Auto generated"
          key={form.key('id')}
          {...form.getInputProps('id')}
        />

        <Select
          disabled={mode == WatchedChannelEditMode.Edit}
          label="Channel"
          data={channelSelect}
          key={form.key('channel_id')}
          {...form.getInputProps('channel_id')}
          searchable
        />

        <Select
          label="Resolution"
          data={qualityOptions}
          key={form.key('resolution')}
          {...form.getInputProps('resolution')}
          searchable
        />

        <Checkbox
          mt={10}
          label="Archive Chat"
          key={form.key('archive_chat')}
          {...form.getInputProps('archive_chat', { type: "checkbox" })}
        />

        <Checkbox
          mt={5}
          label="Render Chat"
          key={form.key('render_chat')}
          {...form.getInputProps('render_chat', { type: "checkbox" })}
        />

        <Divider my="sm" size="md" />

        <div>
          <Title order={3}>Live Streams</Title>
          <Text>Archive live streams as they are broadcasted.</Text>

          <Checkbox
            mt={5}
            label="Watch Live Streams"
            key={form.key('watch_live')}
            {...form.getInputProps('watch_live', { type: "checkbox" })}
          />
        </div>

        <Divider my="sm" size="md" />

        <div>
          <Title order={3}>Channel Videos</Title>
          <Text>Archive past channel videos.</Text>
          <Text size="xs" fs="italic">
            Check for new videos occurs once a day.
          </Text>

          <Checkbox
            mt={5}
            label="Watch Videos"
            key={form.key('watch_vod')}
            {...form.getInputProps('watch_vod', { type: "checkbox" })}
          />

          <Box ml={30}>
            <Checkbox
              mt={5}
              label="Download Archives"
              description="Download past live streams"
              key={form.key('download_archives')}
              {...form.getInputProps('download_archives', { type: "checkbox" })}
            />
            <Checkbox
              mt={5}
              label="Download Highlights"
              description="Download past highlights"
              key={form.key('download_highlights')}
              {...form.getInputProps('download_highlights', { type: "checkbox" })}
            />
            <Checkbox
              mt={5}
              label="Download Uploads"
              description="Download past uploads"
              key={form.key('download_uploads')}
              {...form.getInputProps('download_uploads', { type: "checkbox" })}
            />
            <Checkbox
              mt={5}
              label="Download Subscriber-only Videos"
              description="Do not check this if you are not a subscriber. Must have Twitch token set in Admin > Settings to download subscriber-only videos."
              key={form.key('download_sub_only')}
              {...form.getInputProps('download_sub_only', { type: "checkbox" })}
            />
          </Box>
        </div>

        <Divider my="sm" size="md" />

        <div>
          <Title order={3}>Channel Clips</Title>
          <Text>Archive past channel clips.</Text>
          <Text size="xs">
            This feature is meant to archive the most popular (view count) clips the past <code>interval</code> days. For example, if you set number of clips to 5 and interval to 7, the top 5 clips from the past 7 days will be archived. It will not run again until 7 days have passed. No restrictions (categories, age, title, regex, etc) are applied to this.
          </Text>

          <Checkbox
            my={5}
            label="Watch Clips"
            key={form.key('watch_clips')}
            {...form.getInputProps('watch_clips', { type: "checkbox" })}
          />


          <NumberInput
            label="Number of Clips"
            description="Number of clips to archive."
            key={form.key('clips_limit')}
            {...form.getInputProps('clips_limit')}
            min={1}
          />

          <NumberInput
            label="Interval Days"
            description="How often channel is checked for clips to archive (in days). This also limits the clip search to the last interval days."
            key={form.key('clips_interval_days')}
            {...form.getInputProps('clips_interval_days')}
            min={1}
          />

        </div>

        <Divider my="sm" size="md" />

        <div>
          <Title order={3}>Advanced</Title>

          <NumberInput
            label="Max Video Age (days)"
            description="Archive videos that are not older than this number of days (0 to archive all)."
            key={form.key('video_age')}
            {...form.getInputProps('video_age')}
          />

          <Group mt={5}>
            <Title order={5}>Title Regex</Title>
            <Tooltip label="Add title regex">
              <ActionIcon size="sm" variant="filled" color="green" aria-label="Add Title Regex" onClick={() => {
                const newRegex: WatchedChannelTitleRegex = {
                  apply_to_videos: false,
                  negative: false,
                  regex: "",
                  id: ""
                }
                setLiveTitleRegexes(liveTitleRegexes => [...(liveTitleRegexes ?? []), newRegex])
              }}>
                <IconPlus style={{ width: '70%', height: '70%' }} stroke={1.5} />
              </ActionIcon>
            </Tooltip>
          </Group>
          <div>
            <Text size="sm">Use regex to filter and match specific patterns in livestream and video titles. See <a className={classes.link} href="https://github.com/Zibbp/ganymede/wiki/Watched-Channel-Title-Regex" target="_blank">wiki</a> for more information.</Text>
          </div>

          <div>
            {liveTitleRegexes && liveTitleRegexes.map((regex: WatchedChannelTitleRegex, index) => (
              <div key={index}>
                <Grid grow>
                  <Grid.Col span={10}>
                    <TextInput
                      label="Regex"
                      placeholder="(?i:rerun)"
                      value={regex.regex}  // Use the specific regex from the current item
                      onChange={(e) => {
                        const updatedRegexes = [...liveTitleRegexes];
                        updatedRegexes[index] = {
                          ...updatedRegexes[index],
                          regex: e.currentTarget.value,
                        };
                        setLiveTitleRegexes(updatedRegexes);
                      }}
                    />

                    <Group mt={7}>
                      <Checkbox
                        defaultChecked
                        label="Negative"
                        description="Invert match"
                        checked={regex.negative}
                        onChange={(e) => {
                          const updatedRegexes = [...liveTitleRegexes];
                          updatedRegexes[index].negative = e.currentTarget.checked;
                          setLiveTitleRegexes(updatedRegexes)
                        }}
                      />
                      <Checkbox
                        defaultChecked
                        label="Apply to video downloads"
                        description="Applies to live streams only by default"
                        checked={regex.apply_to_videos}
                        onChange={(e) => {
                          const updatedRegexes = [...liveTitleRegexes];
                          updatedRegexes[index].apply_to_videos = e.currentTarget.checked;
                          setLiveTitleRegexes(updatedRegexes)
                        }}
                      />
                    </Group>
                  </Grid.Col>
                  <Grid.Col span={1} mt={25}>
                    <Group>
                      <ActionIcon size="lg" variant="filled" color="red" aria-label="Settings" h={80} onClick={() => {
                        const updatedRegexs = [...liveTitleRegexes]
                        updatedRegexs.splice(index, 1)
                        setLiveTitleRegexes(updatedRegexs)
                      }}>
                        <IconTrash style={{ width: '70%', height: '70%' }} stroke={1.5} />
                      </ActionIcon>
                    </Group>
                  </Grid.Col>
                </Grid>
              </div>
            ))}
          </div>
        </div>

        <Divider my="sm" size="md" />

        <div>
          <Title order={3}>Categories</Title>
          <Text size="sm">Archive videso and live streams from select categories. Leave blank to archive all categories.</Text>
        </div>

        <Checkbox
          mt={5}
          label="Apply to livestreams"
          description="Apply category restrictions to livestreams"
          key={form.key('apply_categories_to_live')}
          {...form.getInputProps('apply_categories_to_live', { type: "checkbox" })}
        />


        <Box mt={10}>
          {formattedTwitchCategories.length == 0 ? (
            <Button variant="filled" color="violet" onClick={() => handleGetTwitchCategories()}
              loading={twitchCategoriesLoading}
            >Load categories</Button>
          ) : (
            <MultiSelect
              searchable
              limit={20}
              data={formattedTwitchCategories}
              comboboxProps={{ position: 'top', middlewares: { flip: false, shift: false } }}
              placeholder="Search for a category"
              clearable
              key={form.key('categories')}
              {...form.getInputProps('categories')}
            />
          )}
        </Box>

        <Button
          mt={10}
          type="submit"
          fullWidth
          loading={editWatchedChannelMutation.isPending}
        >
          {mode === WatchedChannelEditMode.Create ? 'Create Watched Channel' : 'Save Watched Channel'}
        </Button>
      </form>

    </div>
  );
}

export default AdminWatchedChannelDrawerContent;