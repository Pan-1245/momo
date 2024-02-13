import "dotenv/config";

const getEnvVar = (name) => {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Missing environment variable: ${name}`);
  }
  return value;
};

export const config = {
  DISCORD_TOKEN: getEnvVar("DISCORD_TOKEN"),
  DISCORD_CLIENT_ID: getEnvVar("DISCORD_CLIENT_ID"),
  DISCORD_GUILD_ID: getEnvVar("DISCORD_GUILD_ID"),
};
