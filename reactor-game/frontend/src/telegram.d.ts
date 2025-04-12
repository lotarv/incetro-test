interface TelegramWebApp {
    initData: {
      user?: {
        id: number
        first_name?: string
      }
      hash?: string
      auth_date?: string
    }
    // Можно добавить другие методы/свойства Telegram.WebApp, если понадобятся
  }
  
  interface Window {
    Telegram?: {
      WebApp: TelegramWebApp
    }
  }