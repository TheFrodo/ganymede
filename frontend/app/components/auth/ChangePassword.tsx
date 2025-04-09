import { authChangePassword } from "@/app/hooks/useAuthentication";
import { Button, PasswordInput } from "@mantine/core";
import { useForm, zodResolver } from "@mantine/form";
import { showNotification } from "@mantine/notifications";
import { useState } from "react";
import { z } from "zod";

const schema = z.object({
  password: z.string().min(8, { message: "Das Passwort sollte mindestens 8 Zeichen haben" }),
  new_password: z.string().min(8, { message: "Das Passwort sollte mindestens 8 Zeichen enthalten" }),
  confirm_new_password: z.string().min(8, { message: "Das Passwort sollte mindestens 8 Zeichen enthalten" })
})

type Props = {
  handleClose: () => void;
}

const AuthChangePassword = ({ handleClose }: Props) => {
  const [loading, setLoading] = useState(false)
  const form = useForm({
    mode: "uncontrolled",
    initialValues: {
      password: '',
      new_password: '',
      confirm_new_password: '',
    },

    validate: zodResolver(schema),
  });

  const handleSubmit = async (password: string, newPassword: string, confirmNewPassword: string) => {
    if (newPassword != confirmNewPassword) {
      showNotification({
        message: "Neue Passwörter müssen übereinstimmen",
        color: "red"
      })
      return
    }
    try {
      setLoading(true)
      await authChangePassword(password, newPassword, confirmNewPassword)
      showNotification({
        message: "Passwort geändert"
      })
      handleClose()
    } catch (error) {
      console.error(error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div>
      <form onSubmit={form.onSubmit((values) => handleSubmit(values.password, values.new_password, values.confirm_new_password))}>
        <PasswordInput
          label="Aktuelles Passwort"
          key={form.key('password')}
          {...form.getInputProps('password')}
          radius="md"
        />
        <PasswordInput
          label="Neues Passwort"
          key={form.key('new_password')}
          {...form.getInputProps('new_password')}
          radius="md"
        />
        <PasswordInput
          label="Bestätige das neue Passwort"
          key={form.key('confirm_new_password')}
          {...form.getInputProps('confirm_new_password')}
          radius="md"
        />

        <Button mt={10} type="submit" loading={loading} fullWidth>
          Passwort ändern
        </Button>
      </form>
    </div>
  );
}

export default AuthChangePassword;