import http from 'k6/http';
import { sleep, check } from 'k6';

// Configuração do teste
export const options = {
  vus: 100,           // 100 usuários virtuais simultâneos
  duration: '60s',    // 1 minuto de duração
};

// Função principal DEVE ser exportada como default
export default function () {  // ⬅️ Garanta que está escrito exatamente assim
  const url = 'http://127.0.0.1:63342/health'; 

  const res = http.get(url);

  check(res, {
    'Status 200': (r) => r.status === 200,
  });

  sleep(1);
}