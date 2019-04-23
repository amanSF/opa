package util

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func DecodeStringAndWriteFile(s string, path string) error {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, decoded, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Test_KeypairReloader(t *testing.T) {

	var pemCert1 = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURBakNDQWVxZ0F3SUJBZ0lKQU5DaTNHVUoycUthTUEwR0NTcUdTSWIzRFFFQkJRVUFNQmN4RlRBVEJnTlYKQkFNTURHRmtiV2x6YzJsdmJsOWpZVEFnRncweE9EQTBNakF4TnpBMU1EZGFHQTh5TWpreU1ESXdNekUzTURVdwpOMW93TFRFck1Da0dBMVVFQXd3aVlXUnRhWE56YVc5dUxXTnZiblJ5YjJ4c1pYSXViM0JoTFdsemRHbHZMbk4yCll6Q0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1TejhoYkNkaENJVXhLQ3Z3ODQKbFZES1YxWkc4cnFwcDNyZ1AydUF1cHRYS0JzcGcwU2ZSZlR5aVBuQzhMaTVlRy9YbXJEdmhxNGtkSjJ1a05VKwpTZjhtbGhzRDlHSXpkUFlNUnJCQjJvbnp5eXZrYU1oYUpIai8remUxSlNqUTZZWjgyKzNGMndST2RZbE16UzZGCmt4UGdJUEc2dXVTdzBKNmFPZEtXZTFJZzhRS04rM3J4SGd6QzhXRWNaS0xjT0tiQ00xMWJVQXpLMm1PaXgwQkkKL1NiTlY1UElqSlg3ZWtzZldXdFZjVCt5cWVDV3FOSC95ME5jWko3aXlYUHROKytRQjRYUEx0YnBuZC9iS2N1TwpLKzkxSUZ6dVhxYzUwMEQ4REN3OTJJNkJqMnpNTGFBd2hCSlA2L1c5dUZqUkF5UTlDNEtZcTNzdlFhMm0vdDFuClRWc0NBd0VBQWFNNU1EY3dDUVlEVlIwVEJBSXdBREFMQmdOVkhROEVCQU1DQmVBd0hRWURWUjBsQkJZd0ZBWUkKS3dZQkJRVUhBd0lHQ0NzR0FRVUZCd01CTUEwR0NTcUdTSWIzRFFFQkJRVUFBNElCQVFCSVovb1p5Y0EwSFZzcQp1anBPaDdDaVdwaWIwWHVFNlo1TVNuS00wSUxnVEtuYndHc0RLd0NjK3N0RmF5blY2WkhVL1NLaHNnRlk1b3NjCloya3Z1VWpQb2dZUmlRYXZqWHJPZWJnd0s0b0hnSnhGUTZTTkI0bVNVb2VDcERIbjkxZnc0SVg0RmRGcXk3WGgKVVpqNHUxbS92UllISzFEU1ZRd3hEU2lldnFEQjcyTy9waktmais1Nm1yeW9GMTArWHRxNlpKTHVpUWROc3FZZgpoMndoRTFGa25lNldkOXBIUVVtblpWRXpXbyt1WVB2WDhqdyt5ZUR4K0cyY0tnUTF6U0tEUE16RWNPVDR2bDFMClE3Ym40Z04zTXkzS3hwTnFRVnFqT0FTVytWNWRvUWtnUWw3OEV5QkNWcy9lQUFrdFl4aVE2U3NJRTRNL0NhMlIKUXpIQ1hTT3UKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
	var pemKey1 = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb2dJQkFBS0NBUUVBeExQeUZzSjJFSWhURW9LL0R6aVZVTXBYVmtieXVxbW5ldUEvYTRDNm0xY29HeW1EClJKOUY5UEtJK2NMd3VMbDRiOWVhc08rR3JpUjBuYTZRMVQ1Si95YVdHd1AwWWpOMDlneEdzRUhhaWZQTEsrUm8KeUZva2VQLzdON1VsS05EcGhuemI3Y1hiQkU1MWlVek5Mb1dURStBZzhicTY1TERRbnBvNTBwWjdVaUR4QW8zNwpldkVlRE1MeFlSeGtvdHc0cHNJelhWdFFETXJhWTZMSFFFajlKczFYazhpTWxmdDZTeDlaYTFWeFA3S3A0SmFvCjBmL0xRMXhrbnVMSmMrMDM3NUFIaGM4dTF1bWQzOXNweTQ0cjczVWdYTzVlcHpuVFFQd01MRDNZam9HUGJNd3QKb0RDRUVrL3I5YjI0V05FREpEMExncGlyZXk5QnJhYiszV2ROV3dJREFRQUJBb0lCQUM4dlgwNVlYWXYvdDdYRQpDV0wwVnRVbmhLWGNwUk5qNlJvUHV3Ymx2VjNRSFgzT2luK21tVmgxNmFmSVBJaURiVUdlL2F1ZCtiZkNaUHNDClNobnRORUR4bDZacFgzRkRHTGFCSVYvY25EamxjNEpXNjAvTVY5MXl3bWVObTVsYjYvam44TFhyOVdybUMwbmgKVTRMcVplcXBwQjE0aHFhZlI1d1VNUFMwUm1iZnRrNVluU1ZDMkYrMDYwZXZjcEVrL3ZGT2Fxd2JQWE5vL200ZAp1RDdZVllqcmdBMzZySFNOOXQwbTNwakc3dWFleDYyeU5PKzF5ejBGdjFrSnIzaEh6cG1Vb3FlQlNnYTJJR2orCml2TTJjaHNGdVN1YWNZV1Fyc3JuZ0NRNi9ySUdZTjZGRXQ2NVV1T2dvTi9OWUZTVGt0SEtBTEgvZkcvMWpkL0sKdm03dGdRRUNnWUVBNmp5NnFBRlBEWExiNEY0cjMzR2xMbThRZi9GZDlxdTdvbTQvV29yWmhwMWpzdHJQNWc3NQphMjFGM05uayt1bGo3SFliby83QnZJNnVSTjNmamw0S1V6d0ZyR2RMN3g3bXhmMFRvNUFGZUs1L3dZRFM4N1M3CkIyRXc1VFFnQ0hiK1QxbnJKMEU4UG9TcUpqM3JKUEJTYktmdnFzL0hiekQwUC9zbHZNZWNwZHNDZ1lFQTF2cDUKQ2pNSzAxRFF3ZldYVkRkOXp2em5PZ0FMZFUvUzQ1OUd3azNrZ1lJTlA2dzkzZEx1NnVqaytJc3BUbjloZ3F3Rwpld2dlN0dIWURKZGJsRW9lNXRRRjZpdmgrdk9meVRwejlMMUxXbmROQWZSY05hWlUrUzdRbHg4RVA0NGZzU0tICllCbEp6U0pnY1lHNUkxdnN3ak82UXJoQTBvUzQrMlBIT2dOOVRvRUNnWUFIV0VZbUZIeHRHSWZxZmlsMEJUVS8Kd24zSzlUK1VCNGlIckZ0U05INWpxVFhDR3ZoZjUySk15dzFnd25oYW9jemZVa2pGbUt4c2hERFV1ZnhUazRGUgpZTlZ3dFlCOURBUUNlMUFOVmV2Ri8vckRqNDJOMU1IMGxraVpOMHcxMTcxYnNaOGRDUCtobWpsWG42TnE4aG96CmpBU0kzNkVLVEllc0plem1BWUJZOXdLQmdENENrZGQzVWJBU3A2VExDcE9vVGduZXVYUHBmMlNmV2QxK25CS0EKRldHbCtkeGIrcHg1czZEZS9PMFVkeDNGY0lNWXkrWEJPZXBGYnVSeTVGK3A0YkFEaUpFN2h6dXorbXM1Q1NtWApVSjBQdkk3THhhMVVCVDY3V2orOUdxU2FnbG05OC82c1RMVjNMUXFRdUs2U1hZREhHdnNUMnQ3ZW1kMHBzdFh4Ckc3d0JBb0dBRU9UYm9iU0hpVFpJYlAyWlJTV1NqQWEyS2VBcmZsVy9iSWlMb1Z2NEpVMXV5SW56WkU0dSsrNEQKRDZuMmFYeHpXSnJNNnlwRGg4SXFqU0hvRTZ0elJFS2pEVEtILzQ0UjVZZFEzTzBQSlVla0V5bk4xa3JJQktFRgpxMWNHMHNCb1h5S2RXMXdhRFFvNzNpNllqa0E0a2VMTzJXZnFuUExCaUlSekE4UG5Qem89Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="
	var pemCert2 = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVLekNDQWhNQ0NRQ0Z3ZjQrdjdFZlREQU5CZ2txaGtpRzl3MEJBUXNGQURCYU1Rc3dDUVlEVlFRR0V3SlYKVXpFTE1Ba0dBMVVFQ0F3Q1EwRXhDekFKQmdOVkJBY01BbE5HTVE0d0RBWURWUVFLREFWTmVVOXlaekVOTUFzRwpBMVVFQ3d3RVRYbFBWVEVTTUJBR0ExVUVBd3dKYkc5allXeG9iM04wTUI0WERURTVNREl5TXpBNE1EUXpOMW9YCkRUSXdNRGN3TnpBNE1EUXpOMW93VlRFTE1Ba0dBMVVFQmhNQ1ZWTXhDekFKQmdOVkJBZ01Ba05CTVFzd0NRWUQKVlFRSERBSlRSakVOTUFzR0ExVUVDZ3dFVFhsUFZURU9NQXdHQTFVRUN3d0ZUWGxQY21jeERUQUxCZ05WQkFNTQpCRTE1UTA0d2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUUNTZDJLTG1DMjZRK2tDCm5PckYvUTVLeTUxMDVPNXd0czNLQUhKdHpUczhCY3ptbUFHSktQUDNLMEc2ZTlUczdRVlZnNStvWGxNRWZmQnUKMWRXZW9rUGw4NTRhVVUvS0dKSFFCQkpiZWRYS0pDdnhoVjhSdG40WnROeTZzZlVBV25vSWUrREdtRzRkTU9UcgpjZXV5WDkyZ09iL0lHNnVKSGNzb2V5Z010VHF1UU1BYXMrYU0zb1V4MSt3MWNGN2RQbEcwMWgzTzU5WFk1bVFVCjNiYkRiNy9WdUFHTzhqNkluaDFSRlBuTHFFeUZLYjZSb0x3Vk5sRkU3WXQvU0pyYlQ2Qm5MR1VyMzc4OVAxUWUKZTJndDF2emJrdlRwdzVyVW90QVlQRGxOY2sxcTBDb2ZCZjRJd2dxVk42MG9LV2pSR2xsODZ6ajE5VVRpbnVKawpPUDRuQzc4UEFnTUJBQUV3RFFZSktvWklodmNOQVFFTEJRQURnZ0lCQUpycDlnTEFlZzI1ZUpnYXdZU3FyNnpxCmNZQmJJNnRoSy9uRjRtemZPb2pwMWtoRVdtMDQxcy94d2N3OWtoRDhIaGptVEV1eDdTbjZCdHlzR0RkRVR3YTkKaVpBeSsvblM1VUEwN2ZjN2RScDk5MlY3dFF2bkdHS1ZXU3lSR3NzaFZuZUZlK0dTem9rVDFkKytKVkZ0NnV2NApJVHpOMjkzNzh6bHN4NUVkN0E1dEQzTU5RVFliMjA3aUtPNjdvaklTcnlYRExyMitqU2dIbTZzMktiZ0VvaHI2CjIwdTROWWNQZkNzVmFqNEZRT2ZERHRwVE5WVlNWOTJtRk42c3dvM1JyeUhNc3prTzZxeXprTS9lNjFRUnpkRUIKWHh2ZlRCWWtZcjNGL0lNVHVwbzM3VnJIL3RBdjhITXpxRDMvdHVzSnRva2RTc1FWS25Uc3VLUHdoRnU2dTFIdgpVQnRPV2NwZ3RybWlBNy9JUlRIa05JM0UvZURkczhWMldmZ0NNNllrdjFxRlRXMVVVOFkycmpxbFJ1SUpJcVdvCmR1T2dNanhKODJRQnZlSHQ3YzVzdWZCcVdnQWN1ZXNOa2hhOUpTU09neUhhYVRzNnF3d2lVa3RsQk9kVFZ0SlMKMk9VWXQwTnhhMll2RmdBbWJnYVh1Z2p1YTdnZE1GWVVQTUo1dlc5WmJEL2tsYWRGQWU5Ynl0ZzdHb2haekVkNQpyUVpWaUdnR3paWTUxYTJGdytHT0pScHVDWGRuOFdHVjVWMHJKSlVDK1lZaSs0aWtNTVFwTHJkS1JqYVloMWRBClJHNm5va29kR1JHeVgzMTUzdHhLQjZxdDlGN2ZMbDczQmVObmpxakkzVjJWOGlWdEJIdGdkci9JZ0JHak4rZWYKZFdrcVVJL1J0dlU2Tm42eHNQNkQKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQ=="
	var pemKey2 = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBa25kaWk1Z3R1a1BwQXB6cXhmME9Tc3VkZE9UdWNMYk55Z0J5YmMwN1BBWE01cGdCCmlTano5eXRCdW52VTdPMEZWWU9mcUY1VEJIM3didFhWbnFKRDVmT2VHbEZQeWhpUjBBUVNXM25WeWlRcjhZVmYKRWJaK0diVGN1ckgxQUZwNkNIdmd4cGh1SFREazYzSHJzbC9kb0RtL3lCdXJpUjNMS0hzb0RMVTZya0RBR3JQbQpqTjZGTWRmc05YQmUzVDVSdE5ZZHp1ZlYyT1prRk4yMncyKy8xYmdCanZJK2lKNGRVUlQ1eTZoTWhTbStrYUM4CkZUWlJSTzJMZjBpYTIwK2daeXhsSzkrL1BUOVVIbnRvTGRiODI1TDA2Y09hMUtMUUdEdzVUWEpOYXRBcUh3WCsKQ01JS2xUZXRLQ2xvMFJwWmZPczQ5ZlZFNHA3aVpEaitKd3UvRHdJREFRQUJBb0lCQUdtOFpoZ0F0SUpUcEZPTwpsSzU2RVdkTnQwNDhOeWFIMGVpU2lnbGFyL2lVOTJkQk5WcWZoZ0JaNmdtR2ZDRE91NXNRZXFKQjlneTRIbXJsClZHSkJJbU5QakVQVGo1YUdSajIxc3NyM0dhL2tqUlMvQS9MMW0vTlM3a2FaVWhYL096b0V0RENxdWxPNGlJcXcKbzBjU2poeXRsQllTWTdPSkhyRUt0ODcweG5na2p2VkRHMkJwN2FUMWZ6Sk5uTEhPVXIxK0xSRVJoNkwzTWJuSQo1Vi9DbVc0MG82Y1M3c1dXUVp4UzVFaTRiN2FjaUZ2TlBXbUFzVHBWRStIVUQ2S0tlVzFTd3NjSzMzdHQ0Rzg5CjNMbHRVREFRVEV1M0pIelYwYjFsa21OM1FOaFlnU2dGSXQwanZjcHg3YlFaOGZxNCtFNzUvcTBKSFBhdXVMbzQKTTVqVmQ0RUNnWUVBd2kvd1FJWmRpbVNocVBKMlNCaWNhZ0owNjlmN2tmb2tPdUNLSVA1NC9MY2RLVU1mZTFmUApubS9sTDFZVkhLckdwVDJkVUNPU1JJblU0SlFUckhLODA2dE1FQkxFdXU4T1YyejNZZ0lzbFRBS2dlS1ZlaW9oCkYvczM0Q1Zqa1BLbjEvUXd6NEphVWVwUDNkY0JVbmZ3bVBMcm16eHNKdVZtdHpyODEyUVJMdHNDZ1lFQXdSYTcKbC9zNGNPVVc1bitZbkoxdzkwMkVCd2Nlc05mQWs4V2dqZGdqNTNuemZLTWtIRkRtQVQ4WWFadEFvdFFza0JBYgozMCtWaUtVKzIwYkNZcnFCS1JELzJFcWpnbExMU09Uek5RM3FsMHpSYzN1RWNPWVp3RklNSkIvSDVTbFRDWWRZCnR6YWljalM3c3crcVF4VDRUVVFJdUJlRFdXRlNrN0hCNjZBMnBOMENnWUE2QkJDSGRwMWp4NFQwOE51ZDFweUUKbjQ5SEZ4UTJITUhPUlY5L1ZBZlptT3ZCZENobnRXaDJNWUFMa0hCbVRDOWxGdTkxRGV1WlJvU3ZSQTVNcTByMwpBbXFMTXRCTCtUZ1RycDZLVFpQOFRvb1ZIQmtFa3Ftb2ZJbkpFZUtaWDZYOXd4c0NKRFBQM1cxWEhmRnJMaUpYCno5dk1XVlVHbWVzQ08wRm5LS0pwbVFLQmdRQ3dlWkIyOVpxK1JBTk1JKzVkcGQvcmh5TldNUUd4NjE4Tk1DRHYKdzAzOWNBNUVLZThGR2tNOWZHZFpqaFhqUER5R1p0S3lwTkFzbGFyL1NvcTdKVzZMRkFoaHJNWUloSmRiUmpXcQpsYytFL1NhcGY0aXFqd09XTE9iLzFUUWpsRU5hZ0NOclFEbjBXNUg1dENXUExnbEt4UldhKzVLSERGWW11dDViCjlwQTYyUUtCZ0dJRWZRTm9EY3A3OWFaWG9RY3AwaXNPVXk0aFZ6WFpHb3V5enNjc2wySFlyODVmRWRSdUxRdnMKNmEyT2FJK01KTU5LSW1mRFNabm9pczh5alMzT3A1OUdET21Gd0J1dFMvbGt0ZFdrcWtTUkVqdmF5OHpGNDM3awpuWDdZdDUrai9NMDVleVo0ZGI1VW00Q0ZScUsxaTdseDdaNWsvUTFzUmd6WHlITFRXZVRQCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg=="

	now := time.Now().UTC()

	pemFile := filepath.Join(os.TempDir(), "test-BuildTLSConfig.pem")
	keyFile := filepath.Join(os.TempDir(), "test-BuildTLSConfig-key.pem")

	err := DecodeStringAndWriteFile(pemCert1, pemFile)
	if err != nil {
		t.Fatalf("Failed to load cert1: %v", err)
	}
	err = DecodeStringAndWriteFile(pemKey1, keyFile)
	if err != nil {
		t.Fatalf("Failed to load key1: %v", err)
	}

	time.Sleep(100)

	k, err := NewKeypairReloader(pemFile, keyFile, 100*time.Millisecond)
	if err != nil || k == nil {
		t.Fatalf("Unexpected error calling NewKeypairReloader: %v", err)
	}
	defer k.Close()

	loadedAt := k.LoadedAt()
	if !loadedAt.After(now) || k.LoadedCount() != 1 {
		t.Fatalf("loaded time must be after test start time, and count==1")
	}

	err = DecodeStringAndWriteFile(pemCert2, pemFile)
	if err != nil {
		t.Fatalf("Failed to load cert2: %v", err)
	}
	err = DecodeStringAndWriteFile(pemKey2, keyFile)
	if err != nil {
		t.Fatalf("Failed to load key2: %v", err)
	}

	time.Sleep(200 * time.Millisecond)

	loadedAt2 := k.LoadedAt()
	if !loadedAt2.After(loadedAt) || k.LoadedCount() != 2 {
		t.Fatalf("loaded time must be after test start time, and count==2")
	}
}
