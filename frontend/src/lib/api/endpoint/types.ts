import { z } from 'zod'

export const developer = z.object({
  id: z.number().nonnegative('Developer ID should be non-negative'),
  name: z.string().optional(),
  login: z.string().min(1, 'Login should not be empty'),
  avatar_url: z.string().url('Avatar URL should be a valid URL'),
  company: z.string().nullable(),
  location: z.string().nullable(),
  bio: z.string().nullable(),
  blog: z.string().nullable(),
  email: z.string().email('Email should be a valid email').or(z.literal('')).or(z.null()),

  followers: z.number().nonnegative('Followers count should be non-negative'),
  following: z.number().nonnegative('Following count should be non-negative'),

  stars: z.number().nonnegative('Stars count should be non-negative'),
  repos: z.number().nonnegative('Repos count should be non-negative'),
  gists: z.number().nonnegative('Gists count should be non-negative'),

  created_at: z.coerce.date(),
  updated_at: z.coerce.date(),
})
export type Developer = z.infer<typeof developer>

export const developerPulsePoint = z.object({
  pulse_point: z.object({
    id: z.number().nonnegative('Developer ID should be non-negative'),
    pulse_point: z.number().nonnegative('Pulse Point should be non-negative'),
    updated_at: z.coerce.date(),
  }),
})
export type DeveloperPulsePoint = z.infer<typeof developerPulsePoint>

export const developerWithPulsePoint = z.object({
  developer: developer,
  pulse_point: developerPulsePoint,
})
export type DeveloperWithPulsePoint = z.infer<typeof developerWithPulsePoint>

export const language = z.object({
  id: z.string().min(1, 'Language ID should not be empty'),
  name: z.string().min(1, 'Language name should not be empty'),
  color: z.string().min(1, 'Language color should not be empty'),
})
export type Language = z.infer<typeof language>

export const languageWithUsage = z.object({
  language: language,
  // 0-100
  percentage: z
    .number()
    .min(0, 'Language usage percentage should be non-negative')
    .max(100, 'Language usage percentage should be less than 100'),
})
export type LanguageWithUsage = z.infer<typeof languageWithUsage>

export const developerLanguages = z.object({
  languages: z.object({
    id: z.number().nonnegative('Developer ID should be non-negative'),
    languages: z.array(languageWithUsage),
    updated_at: z.coerce.date(),
  }),
})
export type DeveloperLanguages = z.infer<typeof developerLanguages>

export const developerRegion = z.object({
  region: z.object({
    id: z.number().nonnegative('Developer ID should be non-negative'),
    region: z.string().min(1, 'Region should not be empty'),
    confidence: z
      .number()
      .min(0, 'Region confidence should be non-negative')
      .max(1, 'Region confidence should be less than 1'),
  }),
})
export type DeveloperRegion = z.infer<typeof developerRegion>
