-- 요구사항:
-- 1. HOUSE B/L (HBL_NO) 를 기준으로 각 B/L 별 컨테이너 수량(CNTR 개수) 을 집계
-- 2. 가장 많은 수량을 가진 B/L 1건을 조회
-- 3. 컨테이너 수량은 CNTR_NO 기준으로 COUNT
-- 4. 동일 수량 시 ETD 빠른 순으로 우선 선택
-- 5. 결과 컬럼: HBL_NO, CNTR_COUNT, ETD
-- 6. 정렬: CNTR_COUNT DESC, ETD ASC

-- 여기에 SQL 쿼리를 작성하세요
SELECT HBL_NO, CNTR_COUNT, ETD
FROM (
         SELECT
             T1.HBL_NO,
             COUNT(*) AS CNTR_COUNT,
             T2.ETD
         FROM FMS_HBL_CNTR T1
         LEFT OUTER JOIN FMS_HBL_MST T2 ON T1.HBL_NO = T2.HBL_NO
         GROUP BY T1.HBL_NO
         ORDER BY COUNT(*) DESC, T2.ETD ASC
     )
WHERE ROWNUM = 1;